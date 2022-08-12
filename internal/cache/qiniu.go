package cache

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go-cloud/conf"
	"mime/multipart"
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
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(c, &ret, upToken, fileName, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	return url + ret.Key, nil
}
