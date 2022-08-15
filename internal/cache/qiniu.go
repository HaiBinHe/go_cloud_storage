package cache

import (
	"bufio"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-cloud/conf"
	"mime/multipart"
	"time"
)

//上传文件到七牛云 -- 普通上传
func QiniuUpload(c context.Context, file multipart.File, fileSize int64, fileName string) (string, error) {
	url := conf.QiniuSetting.QiniuServer
	mac := qbox.NewMac(conf.QiniuSetting.AccessKey, conf.QiniuSetting.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: conf.QiniuSetting.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone: &storage.ZoneHuanan,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(c, &ret, upToken, fileName, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	return url + "/" + ret.Key, nil

}
func QiniuUploadByByte(c context.Context, fileData *bufio.Reader, fileSize int64, fileName string) (string, error) {
	url := conf.QiniuSetting.QiniuServer
	mac := qbox.NewMac(conf.QiniuSetting.AccessKey, conf.QiniuSetting.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: conf.QiniuSetting.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone: &storage.ZoneHuanan,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(c, &ret, upToken, fileName, fileData, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	return url + "/" + ret.Key, nil

}

//七牛云下载文件
func QiniuDowload(c context.Context, fileName string) string {
	domain := conf.QiniuSetting.QiniuServer

	mac := qbox.NewMac(conf.QiniuSetting.AccessKey, conf.QiniuSetting.SecretKey)

	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	accessUrl := storage.MakePrivateURLv2(mac, domain, fileName, deadline)
	return accessUrl
}
