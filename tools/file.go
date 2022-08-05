package tools

import (
	"go-cloud/conf"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeVideo
	TypeFile
)

//上传图片使用
func Type(ext string) FileType {
	ext = strings.ToUpper(ext)
	switch ext {
	case ".JPG", ".JPEG", ".PNG":
		return TypeImage
	case ".AVI", ".WMV", ".MP4", ".MPEG", ".MPG", ".M4V":
		return TypeVideo
	default:
		return TypeFile
	}

}

//将正常的文件名转为md5形式的文件名
func GetMD5FileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	return EncodeMD5(fileName) + ext
}

//获取文件名的扩展名 包括'.'
func GetFileExt(name string) string {
	return strings.ToUpper(path.Ext(name))
}

//获取文件保存路径
func GetSavePath() string {
	return conf.AppSetting.UploadSavePath
}

//检查保存目录是否存在 不存在：true
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//TODO 检查文件是否达到分块标准
func CheckMaxSize(f multipart.File) bool {

	return false
}

//检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//创建在上传文件时所使用的保存目录
func CreatSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

//保存所上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()
	out, err := os.Create(dst)
	if err != nil {
		return nil
	}
	defer func() { _ = out.Close() }()
	_, err = io.Copy(out, src)
	return err
}
