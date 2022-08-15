package mq

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-cloud/internal/cache"
)

var Qiniu = "qiniu"

type TransferMsg struct {
	Topic    string //消息队列名称
	fileName string
	CurPath  string //当前文件所在位置
	DestPath string //传输目的位置 与七牛云

}

//发送消息
func (m *TransferMsg) SendMsg(c context.Context) (string, error) {
	return cache.XAdd(c, &redis.XAddArgs{
		Stream: m.Topic,
		ID:     "*",
		Values: []interface{}{
			"fileName", m.fileName,
			"curPath", m.CurPath,
			"destPath", m.DestPath,
		},
	})
}
