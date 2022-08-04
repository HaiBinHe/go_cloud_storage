package main

import (
	"go-cloud/cmd"
	"go-cloud/internal/db"
	"log"
)

func main() {

}

func init(){
	//初始化配置文件
	err := cmd.InitSettings()
	if err != nil {
		log.Println(err)
		return
	}
	//初始化数据库
	err = db.InitMySQLConn()
	if err != nil {
		log.Println(err)
		return
	}
	err = db.InitRedisConn()
	if err != nil {
		log.Println(err)
		return
	}
}