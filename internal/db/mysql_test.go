package db

import (
	"go-cloud/cmd"
	"testing"
)

func Test_initMySQLConn(t *testing.T) {
	err := cmd.InitSettings()
	if err != nil {
		return
	}
	err = InitMySQLConn()
	if err != nil {
		return
	}
}
