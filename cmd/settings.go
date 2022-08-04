package cmd

import (
	"go-cloud/tools"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	ServerSetting *Server
	AppSetting *AppSettingS
	JWTSetting *JWTSettingS
	DataBaseSetting *DatabaseSettingS
	RedisSetting *RedisSettingS

)

func InitSettings() error{
	ServerSetting = &Server{}
	AppSetting = &AppSettingS{}
	JWTSetting = &JWTSettingS{}
	DataBaseSetting = &DatabaseSettingS{}
	RedisSetting = &RedisSettingS{}
	//获取指定目录所在的根目录
	path := tools.InferRootDir("/conf")
	cfg, err := ini.Load(path + "/conf/config.ini")

	if err != nil {
		log.Fatal("配置文件加载失败:", err)
		return err
	}
	err = cfg.Section("Server").MapTo(ServerSetting)
	if err != nil {
		log.Fatal("Server配置加载失败:", err)
		return err
	}
	err = cfg.Section("App").MapTo(AppSetting)
	if err != nil {
		log.Fatal("App配置加载失败:", err)
		return err
	}
	err = cfg.Section("DataBase").MapTo(DataBaseSetting)
	if err != nil {
		log.Fatal("DataBase配置加载失败:", err)
		return err
	}
	err = cfg.Section("Redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatal("Redis配置加载失败:", err)
		return err
	}
	err = cfg.Section("JWT").MapTo(JWTSetting)
	if err != nil {
		log.Fatal("JWT配置加载失败:", err)
		return err
	}

	ServerSetting.ReadTimeOut *= time.Second
	ServerSetting.WriteTimeOut *= time.Second
	JWTSetting.Expire *= time.Second
	return nil
}
//TODO 通过命令行解析 -port  -configPath -runMode
func InitFlag() error {
	return nil
}
