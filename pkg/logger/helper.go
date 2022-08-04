package logger

import (
	"context"
	"fmt"
	"go-cloud/cmd"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)


type Helper struct {
	L *Logger
}
//默认打印到控制台的日志
func NewStdLogger() *Helper {
	return &Helper{
		L: NewLogger(os.Stdout, "", log.LstdFlags),
	}
}
func InitLogger() (*Helper, error) {
	return &Helper{
		L: NewLogger(&lumberjack.Logger{
			Filename: cmd.AppSetting.LogSavePath + "/" + cmd.AppSetting.LogFileName+ cmd.AppSetting.LogFileExt,
			//文件最长存放10天
			MaxAge: 10,
			//单个文件最大100MB
			MaxSize: 100,
			LocalTime: true,
		}, "", log.LstdFlags),
	},nil
}
func (h *Helper) Info(ctx context.Context, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelInfo, fmt.Sprint(v...))
}
func (h *Helper) Infof(ctx context.Context, format string, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (h *Helper) Fatal(ctx context.Context, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelFatal, fmt.Sprint(v...))
}
func (h *Helper) Fatalf(ctx context.Context, format string, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (h *Helper) Debug(ctx context.Context, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelDebug, fmt.Sprint(v ...))
}
func (h *Helper) Debugf(ctx context.Context, format string, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (h *Helper) Warn(ctx context.Context, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelWarn, fmt.Sprint(v ...))
}
func (h *Helper) Warnf(ctx context.Context, format string, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelWarn, fmt.Sprintf(format, v...))
}

func (h *Helper) Error(ctx context.Context, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelError, fmt.Sprint(v ...))
}
func (h *Helper) Errorf(ctx context.Context, format string, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelError, fmt.Sprintf(format, v...))
}

func (h *Helper) Panic(ctx context.Context, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelPanic, fmt.Sprint(v ...))
}
func (h *Helper) Panicf(ctx context.Context, format string, v ...interface{}) {
	h.L.WithContext(ctx).Output(LevelPanic, fmt.Sprintf(format, v...))
}

func (h *Helper) WithFields(f Fields) *Helper {
	h.L = h.L.WithFields(f)
	return h
}
func (h *Helper) WithCallersFrames() *Helper {
	h.L = h.L.WithCallersFrames()
	return h
}