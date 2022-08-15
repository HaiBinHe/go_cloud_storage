package cache

import (
	"context"
	"go-cloud/conf"
	"log"
	"os"
	"testing"
)

func TestQiniuUpload(t *testing.T) {
	_ = conf.InitSettings()
	c := context.Background()
	info, err := os.Stat("qiniuTest.txt")
	if err != nil {
		log.Println(err)
		return
	}
	file, err := os.Open(info.Name())
	if err != nil {
		log.Println(err)
		return
	}

	path, err := QiniuUpload(c, file, info.Size(), info.Name())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(path)
}

func Test_DownloadQiniu(t *testing.T) {
	_ = conf.InitSettings()
	path := QiniuDowload(context.Background(), "qiniuTest.txt")
	//rgft2o3y9.hn-bkt.clouddn.com/qiniuTest.txt?e=1660494830&token=vUrJnQXUNLdvShp3KxqIQEqj_eSgLlyQEQ1r2Dyh:NpDqf8uAJt-EwkASnSDGJEhw5qc=
	log.Println("http://" + path)
}
