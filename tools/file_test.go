package tools

import (
	"go-cloud/conf"
	"testing"
)

func Test_file(t *testing.T) {
	_ = conf.InitSettings()
	ext := ".MP4"
	str := "aaa.mp4"
	println(Type(ext))
	println(GetFileExt(str))
}
