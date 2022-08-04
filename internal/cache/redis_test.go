package cache

import (
	"go-cloud/conf"
	"testing"
)

func Test_initRedisConn(t *testing.T) {
	err := conf.InitSettings()

	err = InitRedisConn()
	if err != nil {
		return
	}
}
