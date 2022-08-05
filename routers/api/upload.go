package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/internal/model"
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
	Username    string `json:"username" binding:"max=20"`
	FileStoreID uint64 `json:"file_store_id" binding:"required"`
	ParentID    uint64 `json:"parent_id,omitempty" `
	fileHash    string `binding:"-"`
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
	//query
	fs, err := strconv.ParseUint(c.Query("file_store_id"), 10, 64)
	//defaultQuery
	p, err := strconv.ParseUint(c.DefaultQuery("parent_id", "0"), 10, 64)
	if err != nil {
		logger.StdLog().Errorf(c, "parseUint failed:", err)
		response.RespError(c, error2.ServerError)
		return
	}
	uf.FileStoreID = fs
	uf.ParentID = p
	//普通上传
	info, err := upload(file, fileHeader, *uf)
	if err != nil {
		return
	}
	response.RespData(c, info)

}

//TODO 普通上传,分块上传,保存到数据库,OSS存储,Ceph存储
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

	//检查文件是否达到分块大小
	if tools.CheckMaxSize(file) {
		//TODO 分块上传逻辑
		chunksUpload()
	} else {
		//检查是否存在文件目录
		if tools.CheckSavePath(savePath) {
			if err := tools.CreatSavePath(savePath, os.ModePerm); err != nil {
				return nil, error2.DirError
			}
		}
		//检查文件权限
		if tools.CheckPermission(savePath) {
			return nil, error2.DirPermissionError
		}
		//保存文件
		if err := tools.SaveFile(fileHeader, dst); err != nil {
			return nil, err
		}
		//TODO 保存到数据库中

	}
	return &FileInfo{
		Name:       u.fileName,
		Size:       u.fileSize,
		AccessUrl:  accessUrl,
		UploadPath: savePath,
	}, nil
}

//chunksUpload 分块上传
func chunksUpload() {

}

//saveToDB 保存到数据库中
func saveToDB(fileName string, u userUpload) error {
	uf := model.UserFile{
		FileName:     fileName,
		FileHash:     u.fileHash,
		FileSize:     u.fileSize,
		FileStoreID:  u.FileStoreID,
		FileFolderID: u.ParentID,
	}
	err := uf.Create()
	if err != nil {
		logger.StdLog().Error(context.Background(), "save user's FileInfo failed")
		return err
	}
	return nil
}
