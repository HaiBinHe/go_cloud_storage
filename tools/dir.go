package tools

import (
	"os"
	"path/filepath"
)


// InferRootDir 根据已存在的目录名推断出根目录
//path ex. “/conf”
func InferRootDir(path string) string {
	var RootDir string
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + path) {
			return d
		}

		return infer(filepath.Dir(d))
	}

	RootDir = infer(cwd)
	return RootDir
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
