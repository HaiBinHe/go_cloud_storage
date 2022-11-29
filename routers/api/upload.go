package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/cache"
	"go-cloud/internal/dao"
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

// 上传逻辑
func Upload(c *gin.Context) {
	uf := &userUpload{}
	var err error
	valid, _ := app.BindAndValid(c, uf)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}

	formFile, fileHeader, err := c.Request.FormFile("formFile")
	if err != nil {
		logger.StdLog().Errorf(c, "the 'formFile' field does not match :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	// 获取文件基础信息
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
	fileHash, err := tools.FileHash(formFile)
	if err != nil {
		logger.StdLog().Errorf(c, "tools.FileHash failed:", err)
		response.RespError(c, error2.ServerError)
		return
	}
	uf.FileHash = fileHash
	//1.秒传
	ok, info := fastUpload(fileHeader, uf)
	if ok {
		response.RespData(c, info)
		return
	}
	//2.普通上传
	info, err = upload(fileHeader, *uf)
	if err != nil {
		logger.StdLog().Errorf(c, "upload failed:", err)
		response.RespError(c, error2.ServerError)
		return
	}
	//3.保存到七牛云上
	go func() {
		_, err := cache.QiniuUpload(c, formFile, fileHeader.Size, fileHeader.Filename)
		if err != nil {
			logger.StdLog().Errorf(c, "Qiniu upload failed:", err)
			response.RespError(c, error2.ServerError)
			return
		}
	}()
	response.RespData(c, info)
}

//上传文件
func upload(fileHeader *multipart.FileHeader, u userUpload) (*FileInfo, error) {
	//文件保存目录
	savePath := tools.GetSavePath()
	//文件路径
	dst := savePath + "/" + u.Username + "/" + u.fileName

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
	// 保存到用户文件表中
	err = upload2.SaveUserFile(uf)
	if err != nil {
		return nil, err
	}
	// 保存到中心文件表中
	err = upload2.SaveCenterFile(uf)
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		Name:       u.fileName,
		Size:       u.fileSize,
	}, nil
}

//秒传
func fastUpload(fileHeader *multipart.FileHeader, u *userUpload) (bool, *FileInfo) {
	//1.查询中心文件表中是否有相同Hash的文件
	//中心文件表存在相同Hash的文件
	if dao.FileIsExist(u.FileHash) {
		//文件路径
		dst := "center" + u.FileHash
		//2.返回成功以及文件相关信息
		return true, &FileInfo{
			AccessUrl:  "",
			UploadPath: dst,
			Name:       fileHeader.Filename,
			Size:       fileHeader.Size,
		}
	}
	return false, nil
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
