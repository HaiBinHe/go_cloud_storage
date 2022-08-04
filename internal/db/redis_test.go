package db

import (
	"go-cloud/cmd"
	"testing"
)

func Test_initRedisConn(t *testing.T) {
	err := cmd.InitSettings()

	err = InitRedisConn()
	if err != nil {
		return
	}
}
