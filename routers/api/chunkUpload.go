package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/cache"
	upload2 "go-cloud/internal/service/upload"
	"go-cloud/pkg/app"
	error2 "go-cloud/pkg/error"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
	"go-cloud/tools"
	"io/ioutil"
	"os"
	"strconv"
)

type chunkFileInfo struct {
	UserName    string `json:"user_name" binding:"max=20"`
	FileStoreID uint64 `json:"file_store_id" binding:"required"`
	ParentID    uint64 `json:"parent_id,omitempty" `
	FileName    string `json:"file_name" binding:"-"`
	FileExt     string `json:"file_ext" binding:"-"`
	//单个分块文件的Hash值
	FileHash    string `json:"file_hash" binding:"required"`
	UploadID    string `json:"upload_id" binding:"required"`
	ChunkNumber int    `json:"chunk_number" binding:"required"`
	TotalChunk  int    `json:"total_chunk" binding:"required"`
	ChunkSize   int    `json:"chunk_size" binding:"required"`
}

//分块上传
func ChunkUpload(c *gin.Context) {
	var err error
	cf := &chunkFileInfo{}
	valid, err := app.BindAndValid(c, &cf)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	//获取文件
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		logger.StdLog().Errorf(c, "the 'file' field does not match :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	cf.FileName = fileHeader.Filename
	cf.FileExt = tools.GetFileExt(fileHeader.Filename)
	//TODO 秒传判断
	//获取文件夹
	savePath := tools.GetSavePath()
	//文件路径
	dst := savePath + "/" + cf.UserName + "/chunk/" + cf.FileName + "/" + cf.FileHash + "_" + strconv.Itoa(cf.ChunkNumber)
	//1.检查文件夹是否存在以及权限问题
	err = checkDirAndPermission(savePath)
	if err != nil {
		response.RespError(c, error2.InvalidParams)
		return
	}
	//2.保存分块文件到指定目录
	if err := tools.SaveFile(fileHeader, dst); err != nil {
		response.RespError(c, error2.ServerError)
		return
	}
	//3.存入redis中
	data := []byte{}
	_, err = file.Read(data)
	if err != nil {

	}
	err = cache.HSet(c, cf.UploadID,
		[]string{
			"fileHash", cf.FileHash,
			"chunkNumber", strconv.Itoa(cf.ChunkNumber),
			"totalChunk", strconv.Itoa(cf.TotalChunk),
			"chunkSize", strconv.Itoa(cf.ChunkSize),
			"data", string(data),
		})
	if err != nil {
		response.RespError(c, "存入redis失败")
		return
	}
	//4.获取文件夹下面有多少个文件
	chunkList := []string{}
	dirpath := savePath + "/" + cf.UserName + "/chunk/" + cf.FileHash
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		response.RespError(c, "文件读取失败")
		return
	}
	for _, f := range files {
		fileName := f.Name()
		chunkList = append(chunkList, fileName)
	}
	//5.返回数据
	response.RespData(c, map[string]interface{}{
		"chunkList":       chunkList,
		"chunk_file_info": cf,
	})
	return
}

//检查分块文件状态
func CheckChunk(c *gin.Context) {
	uploadID := c.Query("upload_id")
	result, err := cache.HGetAll(c, uploadID)
	//0 分块文件不存在 1:分块文件存在
	if err != nil {
		response.RespData(c, map[string]interface{}{
			"state": 0,
			"msg":   "upload_id:" + uploadID + " 对应的文件未上传.",
		})
		return
	}
	response.RespData(c, map[string]interface{}{
		"state":  1,
		"result": result,
	})
	return

}

//合并分块
func MergeChunk(c *gin.Context) {
	var err error
	cf := &chunkFileInfo{}
	valid, err := app.BindAndValid(c, &cf)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	//文件保存目录
	savePath := tools.GetSavePath()
	//文件路径
	dst := savePath + "/" + cf.UserName + "/" + cf.FileName
	cf.FileExt = tools.GetFileExt(cf.FileName)
	//不存在，可以进行分块上传
	if tools.CheckSavePath(dst) {
		dirpath := savePath + "/" + cf.UserName + "/chunk/" + cf.FileHash
		files, err := ioutil.ReadDir(dirpath)
		if err != nil {
			response.RespError(c, "读取文件夹失败")
			return
		}
		//创建文件
		complateFile, err := os.Create(dst)
		if err != nil {
			response.RespError(c, "创建文件夹失败")
			return
		}
		// 遍历分块文件夹
		var totalSize int64 = 0
		for _, f := range files {
			totalSize += f.Size()
			fileBuffer, err := ioutil.ReadFile(dirpath + "/" + f.Name())
			if err != nil {
				response.RespError(c, "读取文件失败")
				return
			}
			//合并文件
			_, err = complateFile.Write(fileBuffer)
			if err != nil {
				response.RespError(c, "文件写入错误")
				return
			}
		}
		//TODO 记录到数据库中 中心存储数据库和用户文件数据库
		uf := upload2.UserUpload{
			FileStoreID: cf.FileStoreID,
			FileName:    cf.FileName,
			FileHash:    cf.FileHash,
			FileType:    int(tools.Type(cf.FileExt)),
			FileExt:     cf.FileExt,
			FileSize:    totalSize,
			ParentID:    cf.ParentID,
		}
		err = upload2.SaveToDB(uf)
		if err != nil {
			response.RespError(c, "数据库写入错误")
			return
		}
		//TODO 删除redis中的分块缓存

	}
	response.RespData(c, map[string]interface{}{
		"filePath": dst,
	})
	return
}
