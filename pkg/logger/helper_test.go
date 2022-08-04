package logger

import (
	"context"
	"go-cloud/cmd"
	"testing"
)



func TestInitLogger(t *testing.T) {
	err := cmd.InitSettings()
	if err != nil {
		return
	}
	h, err := InitLogger()
	if err != nil {
		return
	}
	ctx := context.Background()
	h.WithFields(Fields{"id":"1"}).Info(ctx, "info msg")
	h.WithFields(Fields{"id":"2"}).Debug(ctx, "debug msg")
	h.WithFields(Fields{"id":"3"}).Warn(ctx, "warn msg")
	h.WithFields(Fields{"id":"4"}).Error(ctx, "error msg")
}
func TestDefaultLogger(t *testing.T){
	err := cmd.InitSettings()
	if err != nil {
		return
	}
	h := NewStdLogger()

	ctx := context.Background()
	h.WithFields(Fields{"id":"1"}).Info(ctx, "info msg")
	h.WithFields(Fields{"id":"2"}).Debug(ctx, "debug msg")
	h.WithFields(Fields{"id":"3"}).Warn(ctx, "warn msg")
	h.WithFields(Fields{"id":"4"}).Error(ctx, "error msg")
}
