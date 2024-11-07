package files

import "mime/multipart"

func IsImage(header *multipart.FileHeader) (string, bool) {
	// 通过文件头部信息判断文件类型
	contentType := header.Header.Get("Content-Type")
	switch contentType {
	case "image/jpeg":
		return contentType, true
	case "image/png":
		return contentType, true
	case "image/gif":
		return contentType, true
	case "image/bmp":
		return contentType, true
	case "application/octet-stream":
		return contentType, true
	case "image/heic":
		return contentType, true
	default:
		return contentType, false
	}
}

func IsVideo() {

}
