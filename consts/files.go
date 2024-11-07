package consts

import (
	"MyIM/pkg/config"
)

// ImageFileType 图片类型,对应相应的路径
type ImageFileType string

const (
	ImageFileTypeAvatar = "avatar"
	ImageFileTypeImage  = "images"
	ImageFileTypeSystem = "system"
	ImageFileTypePhoto  = "photo"
)

// ImageSize 图片大小
type ImageSize string

const (
	ImageSizeOrigin ImageSize = "origin_url_host"
	ImageSizeMiddle           = "middle_url_host"
	ImageSizeThumb            = "thumb_url_host"
)

func GetAvatarPath(fileUrl string) string {
	if fileUrl == "" {
		return config.GetString("uploads.avatar.img_url_host") + "/2c80c1f82101d14311a00beeb7780f61.jpg"
	}
	return GetHostPath(ImageFileTypeAvatar, ImageSizeThumb) + fileUrl
}

func GetSystemPath(fileUrl string) string {
	return GetHostPath(ImageFileTypeSystem, ImageSizeOrigin) + fileUrl
}

func GetFullPath(t ImageFileType, s ImageSize, fileUrl string) string {
	return GetHostPath(t, s) + fileUrl
}

func GetHostPath(t ImageFileType, s ImageSize) string {
	return config.GetString("uploads." + string(t) + "." + string(s))
}
