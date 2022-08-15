package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-cloud/conf"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_initRedisConn(t *testing.T) {
	err := conf.InitSettings()

	err = InitRedisConn()
	if err != nil {
		return
	}
	info, err := os.Stat("qiniuTest.txt")
	if err != nil {
		log.Println(err)
		return
	}

	b, err := ioutil.ReadFile(info.Name())
	if err != nil {
		log.Println(err)
		return
	}
	m := make(map[string]interface{})
	m["fileData"] = b
	ctx := context.Background()
	rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "demo",
		ID:     "*",
		Values: m,
	})
	vals := rdb.XRange(ctx, "demo", "-", "+").Val()
	data := vals[0].Values["fileData"]
	log.Println(data)

}
