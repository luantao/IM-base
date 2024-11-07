package captcha

import "github.com/mojocn/base64Captcha"

// VerifyCaptcha by given id key and remove the captcha value in store, return boolean value.
// 验证图像验证码,返回boolean.
func VerifyCaptcha(identifier, verifyValue string) bool {
	return base64Captcha.VerifyCaptcha(identifier, verifyValue)
}

// VerifyCaptchaAndIsClear verify captcha, return boolean value.
// identifier is the captcha id,
// verifyValue is the captcha image value,
// isClear is whether to clear the value in store.
// 验证图像验证码,返回boolean.
func VerifyCaptchaAndIsClear(identifier, verifyValue string, isClear bool) bool {
	return base64Captcha.VerifyCaptchaAndIsClear(identifier, verifyValue, isClear)
}
