package validate_utils

import "regexp"

// IsvalidateMobile 是否是合格手机号
func IsvalidateMobile(mobileNum string) bool {
	reg := `^1[13456789]{1}\d{9}$`
	regx := regexp.MustCompile(reg)
	return regx.MatchString(mobileNum)
}

// IsvalidateName 汉字/字幕
func IsvalidateName(str string) bool {

	reg := regexp.MustCompile(`^[\p{Han}]+$`)

	if reg.MatchString(str) {
		return true
	}
	return false
}

// IsValidateMenuMark
func IsValidateMenuMark(str string) bool {

	reg := regexp.MustCompile(`^[\w\d]+$`)

	if reg.MatchString(str) {
		return true
	}
	return false
}

// IsPrefixWithSlash
func IsPrefixWithSlash(str string) bool {

	reg := regexp.MustCompile(`^/`)
	regUrl := regexp.MustCompile(`^http`)

	if reg.MatchString(str) || regUrl.MatchString(str) {
		return true
	}
	return false
}
