package cache

import (
	"context"
	"go-cloud/conf"
	"log"
	"testing"
)

func Test_initRedisConn(t *testing.T) {
	err := conf.InitSettings()

	err = InitRedisConn()
	if err != nil {
		return
	}
	str, err := rdb.Get(context.Background(), "demo").Result()
	log.Println(err)
	log.Println(str)
}
