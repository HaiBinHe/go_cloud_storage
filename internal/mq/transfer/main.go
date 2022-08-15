package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"go-cloud/internal/cache"
	"go-cloud/internal/mq"
	"log"
	"os"
)

var (
	groupName string
	consumer  string
	start     string
	count     int64
)

func main() {
	flag.StringVar(&groupName, "group", "g1", "消费者组名")
	flag.StringVar(&consumer, "consumer", "c1", "消费者名")
	flag.StringVar(&start, "start", "0", "开始读取信息的位置")
	flag.Int64Var(&count, "start", 1, "消费者一次性读取信息的个数")
	sg := mq.NewStreamGroup(mq.Qiniu, groupName, consumer, start)
	err := sg.Consume(context.Background(), count, func(msg *mq.ResponseMsg) error {
		//上传文件的信息到七牛云
		curPath := fmt.Sprint(msg.Body["curPath"])
		fileName := fmt.Sprint(msg.Body["fileName"])
		log.Println("当前文件存放于: ", curPath)
		info, err := os.Stat(curPath)
		if err != nil {
			return err
		}

		file, err := os.Open(info.Name())
		if err != nil {
			log.Println("open file err :", err)
			return err
		}
		//上传到七牛云
		path, err := cache.QiniuUploadByByte(context.Background(), bufio.NewReader(file), info.Size(), fileName)
		if err != nil {
			return err
		}
		log.Println("文件访问路径: ", path)
		//TODO 删除临时文件：curPath
		//TODO 保存到数据库也放在这
		return nil
	})
	if err != nil {
		log.Println("消费者处理信息出现错误")
		return
	}
}
