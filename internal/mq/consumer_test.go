package mq

import (
	"context"
	"go-cloud/conf"
	"go-cloud/internal/cache"
	"log"
	"testing"
)

func TestTransferMsg_SendMsg(t *testing.T) {
	_ = conf.InitSettings()
	_ = cache.InitRedisConn()
	m := &TransferMsg{
		Topic:    "demo",
		fileName: "pc2.jpg",
		CurPath:  "./cache",
		DestPath: "七牛云",
	}
	c := context.Background()
	str, err := m.SendMsg(c)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("demoID-->", str)

}
func TestStreamGroup_Consume(t *testing.T) {
	_ = conf.InitSettings()
	_ = cache.InitRedisConn()
	ctx := context.Background()
	c := NewStreamGroup("demo", "g1", "c1", "0")
	err := c.Consume(ctx, 1, func(msg *ResponseMsg) error {
		//将信息上传至七牛云
		curpath := msg.Body["curPath"]
		destPath := msg.Body["destPath"]
		log.Println(curpath)
		log.Println(destPath)
		return nil
	})
	if err != nil {
		return
	}

}
