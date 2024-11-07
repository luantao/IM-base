package files

import (
	"mime/multipart"
	"path/filepath"
)

func GetFileExtension(header *multipart.FileHeader) string {
	// 从文件头部信息中提取文件名及其后缀
	filename := header.Filename
	ext := filepath.Ext(filename)
	if ext == "" {
		return ""
	}
	return ext[1:] // 去除点号（.）并返回后缀字符串
}
