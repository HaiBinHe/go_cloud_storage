package main

import (
	"context"
	"go-cloud/conf"
	"go-cloud/internal/cache"
	"go-cloud/internal/model"
	"go-cloud/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := routers.NewRouters()
	s := &http.Server{
		Addr:           ":" + conf.ServerSetting.Port,
		Handler:        router,
		ReadTimeout:    conf.ServerSetting.ReadTimeOut,
		WriteTimeout:   conf.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	//信号定义
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListeAndServe err :%v", err)
		}
	}()
	//等待中断信号
	quit := make(chan os.Signal)
	//接受 syscall.SIGINT==ctrl+c 和syscall.SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	//最大时间控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server force to shutdown :", err)
	}
	log.Println("Server exiting")
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
