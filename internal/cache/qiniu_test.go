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

	url, err := QiniuUpload(c, file, info.Size(), info.Name())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(url)
}
