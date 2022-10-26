package cache

import (
	"go-cloud/conf"
	"testing"
)

func Test_initMySQLConn(t *testing.T) {
	err := conf.InitSettings()
	if err != nil {
		return
	}
	err = InitMySQLConn()
	if err != nil {
		return
	}
}
