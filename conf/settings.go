package conf

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	ServerSetting   *Server
	AppSetting      *AppSettingS
	JWTSetting      *JWTSettingS
	DataBaseSetting *DatabaseSettingS
	RedisSetting    *RedisSettingS
	QiniuSetting    *QiniuSettingS
)

func InitSettings() error {
	ServerSetting = &Server{}
	AppSetting = &AppSettingS{}
	JWTSetting = &JWTSettingS{}
	DataBaseSetting = &DatabaseSettingS{}
	RedisSetting = &RedisSettingS{}
	QiniuSetting = &QiniuSettingS{}
	//获取指定目录所在的根目录
	path := InferRootDir("/conf")
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
	err = cfg.Section("QINIU").MapTo(QiniuSetting)
	if err != nil {
		log.Fatal("七牛云配置加载失败:", err)
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

// InferRootDir 根据已存在的目录名推断出根目录
//path ex. “/conf”
func InferRootDir(path string) string {
	var RootDir string
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + path) {
			return d
		}

		return infer(filepath.Dir(d))
	}

	RootDir = infer(cwd)
	return RootDir
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
