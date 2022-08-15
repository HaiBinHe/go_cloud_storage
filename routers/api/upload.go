package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/internal/cache"
	"go-cloud/internal/model"
	upload2 "go-cloud/internal/service/upload"
	"go-cloud/pkg/app"
	error2 "go-cloud/pkg/error"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
	"go-cloud/tools"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

type userUpload struct {
	Username    string `form:"username" binding:"max=20"`
	FileStoreID uint64 `form:"file_store_id" binding:"required"`
	ParentID    uint64 `form:"parent_id,omitempty" `
	FileHash    string `form:"file_hash" binding:"required"`
	fileName    string `binding:"-"`
	fileSize    int64  `binding:"-"`
	fileType    int    `binding:"-"`
	fileExt     string `binding:"-"`
}

type FileInfo struct {
	Name       string
	Size       int64
	AccessUrl  string
	UploadPath string
}

func Upload(c *gin.Context) {
	uf := &userUpload{}
	var err error
	valid, _ := app.BindAndValid(c, uf)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		logger.StdLog().Errorf(c, "the 'file' field does not match :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	uf.fileName = strings.Split(fileHeader.Filename, ".")[0]
	uf.fileSize = fileHeader.Size
	uf.fileExt = tools.GetFileExt(fileHeader.Filename)
	uf.fileType = int(tools.Type(uf.fileExt))
	//文件仓库
	fs, err := strconv.ParseUint(c.Query("file_store_id"), 10, 64)
	//父目录
	p, err := strconv.ParseUint(c.DefaultQuery("parent_id", "0"), 10, 64)
	if err != nil {
		logger.StdLog().Errorf(c, "parseUint failed:", err)
		response.RespError(c, error2.ServerError)
		return
	}
	uf.FileStoreID = fs
	uf.ParentID = p
	//1.秒传
	ok, info := fastUpload(fileHeader, uf)
	if ok {
		response.RespData(c, info)
		return
	}
	//TODO 修改上传逻辑
	//2.普通上传
	info, err = upload(file, fileHeader, *uf)
	if err != nil {
		logger.StdLog().Errorf(c, "upload failed:", err)
		response.RespError(c, error2.ServerError)
		return
	}
	//3.保存到七牛云上
	go func() {
		_, err := cache.QiniuUpload(c, file, fileHeader.Size, fileHeader.Filename)
		if err != nil {
			logger.StdLog().Errorf(c, "Qiniu upload failed:", err)
			response.RespError(c, error2.ServerError)
			return
		}
	}()
	response.RespData(c, info)

}

//TODO 普通上传,保存到数据库,OSS存储,Ceph存储
//上传文件
func upload(file multipart.File, fileHeader *multipart.FileHeader, u userUpload) (*FileInfo, error) {
	//文件保存目录
	savePath := tools.GetSavePath()
	//文件路径
	dst := savePath + "/" + u.Username + "/" + u.fileName
	//文件访问路径
	//如果有同名的文件 访问路径不就一样了, 用户名+文件名 生成md5
	md5path := tools.EncodeMD5(u.Username + u.fileName)
	accessUrl := conf.AppSetting.UploadServerUrl + "/" + md5path
	//1.检查文件夹是否存在以及权限问题
	err := checkDirAndPermission(savePath)
	if err != nil {
		return nil, err
	}

	//2.保存文件
	if err := tools.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	//3.文件信息保存到数据库中
	uf := upload2.UserUpload{
		FileStoreID: u.FileStoreID,
		FileName:    u.fileName,
		FileHash:    u.FileHash,
		FileType:    u.fileType,
		FileExt:     u.fileExt,
		FileSize:    u.fileSize,
		ParentID:    u.ParentID,
	}
	err = upload2.SaveToDB(uf)
	if err != nil {
		return nil, err
	}
	//TODO 保存到中心文件表中

	return &FileInfo{
		Name:       u.fileName,
		Size:       u.fileSize,
		AccessUrl:  accessUrl,
		UploadPath: dst,
	}, nil
}

//秒传
func fastUpload(fileHeader *multipart.FileHeader, u *userUpload) (bool, *FileInfo) {
	//1.查询中心文件表中是否有相同Hash的文件
	fc := model.NewFileCenter()
	//中心文件表存在相同Hash的文件
	if fc.FileIsExist(u.FileHash) {
		//文件保存目录
		savePath := tools.GetSavePath()
		//文件路径
		dst := savePath + "/" + u.Username + "/" + u.fileName
		//生成文件路径
		md5path := tools.EncodeMD5(u.Username + u.fileName)
		accessUrl := conf.AppSetting.UploadServerUrl + "/" + md5path

		//2.返回成功以及文件相关信息
		return true, &FileInfo{
			AccessUrl:  accessUrl,
			UploadPath: dst,
			Name:       fileHeader.Filename,
			Size:       fileHeader.Size,
		}
	} else {
		//3.中心文件不存在则进入普通上传模块
		//TODO 同时将数据存入中心文件表中
		return false, nil
	}
}

func checkDirAndPermission(savePath string) error {
	//检查是否存在文件目录
	if tools.CheckSavePath(savePath) {
		if err := tools.CreatSavePath(savePath, os.ModePerm); err != nil {
			return error2.DirError
		}
	}
	//检查文件权限
	if tools.CheckPermission(savePath) {
		return error2.DirPermissionError
	}
	return nil
}
