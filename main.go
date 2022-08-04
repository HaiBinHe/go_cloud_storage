package main

import (
	"go-cloud/conf"
	"go-cloud/internal/cache"
	"go-cloud/internal/model"
	"log"
)

func main() {

}

func init() {
	//初始化配置文件
	err := conf.InitSettings()
	if err != nil {
		log.Println(err)
		return
	}
	//初始化数据库
	err = model.InitMySQLConn()
	if err != nil {
		log.Println(err)
		return
	}
	err = cache.InitRedisConn()
	if err != nil {
		log.Println(err)
		return
	}
}
