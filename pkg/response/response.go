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
		"success":true,
	})
}
func RespError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"success":false,
	})
}

func RespValidatorError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"success":false,
	})
}
func RespData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"success":true,
	})
}
func RespSuccess(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"success":true,
	})

}
