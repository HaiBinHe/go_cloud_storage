package conf

import "time"

type Server struct {
	RunMode      string
	Port         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}
type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	UploadSavePath  string
	UploadServerUrl string
	ShareUrl        string
}
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}
type RedisSettingS struct {
	Host     string
	Port     int
	Password string
	Db       int
	PoolSize int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}
type QiniuSettingS struct {
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
}
