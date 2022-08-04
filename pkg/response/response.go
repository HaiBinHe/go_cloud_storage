package response

import (
	"github.com/gin-gonic/gin"
	"go-cloud/pkg/app"
	"net/http"
)

type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

func RespList(c *gin.Context, list interface{}, totalRows int64) {
	c.JSON(http.StatusOK, gin.H{
		"list": list,
		"page": Pager{
			Page:      app.GetPage(c),
			PageSize:  app.GetPageSize(c),
			TotalRows: totalRows,
		},
	})
}
func RespError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func RespValidatorError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func RespData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}
func RespSuccess(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
