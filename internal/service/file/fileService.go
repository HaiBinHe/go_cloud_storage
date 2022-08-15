package file

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/cache"
)

//文件下载
func DownloadFiles(c *gin.Context) {
	fileDir := c.Query("fileDir")
	fileName := c.Query("fileName")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(fileDir + "/" + fileName)
}

//从七牛云下载文件
func DownloadFilesFromQiniu(c *gin.Context) string {
	fileName := c.Query("fileName")
	return cache.QiniuDowload(c, fileName)

}
