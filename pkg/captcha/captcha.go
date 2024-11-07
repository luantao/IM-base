package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"strings"
)

func Init() {
	base64Captcha.SetCustomStore(NewStore())
	fmt.Printf("\033[1;30;42m[info]\033[0m using redis to captcha %s\n", viper.ConfigFileUsed())
}

// AudioCaptcha 音频验证码
const AudioCaptcha = "audio"

// CharacterCaptcha 字符串验证码
const CharacterCaptcha = "character"

// Digit 数字验证码
const DigitCaptcha = "digit"

// GenerateCaptcha create captcha by config struct and id.
// idkey can be an empty string, base64 will create a unique id four you.
// if idKey is a empty string, the package will generate a random unique identifier for you.
// configuration struct should be one of those struct ConfigAudio, ConfigCharacter, ConfigDigit.
//
// Example OrgCode
//
//	//config struct for digits
//	var configD = base64Captcha.ConfigDigit{
//		Height:     80,
//		Width:      240,
//		MaxSkew:    0.7,
//		DotCount:   80,
//		CaptchaLen: 5,
//	}
//	//config struct for audio
//	var configA = base64Captcha.ConfigAudio{
//		CaptchaLen: 6,
//		Language:   "zh",
//	}
//	//config struct for Character
//	var configC = base64Captcha.ConfigCharacter{
//		Height:             60,
//		Width:              240,
//		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
//		Mode:               base64Captcha.CaptchaModeNumber,
//		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
//		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
//		IsUseSimpleFont:    true,
//		IsShowHollowLine:   false,
//		IsShowNoiseDot:     false,
//		IsShowNoiseText:    false,
//		IsShowSlimeLine:    false,
//		IsShowSineLine:     false,
//		CaptchaLen:         6,
//	}
//	//create a audio captcha.
//	//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
//	idKeyA,capA := base64Captcha.GenerateCaptcha("",configA)
//	//write to base64 string.
//	//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
//	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
//	//create a characters captcha.
//	//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
//	idKeyC,capC := base64Captcha.GenerateCaptcha("",configC)
//	//write to base64 string.
//	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
//	//create a digits captcha.
//	idKeyD,capD := base64Captcha.GenerateCaptcha("",configD)
//	//write to base64 string.
//	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
func GenerateCaptcha(idKey string, configuration interface{}) (id string, captchaInstance base64Captcha.CaptchaInterface) {
	return base64Captcha.GenerateCaptcha(idKey, configuration)
}

// CaptchaWriteToBase64Encoding converts captcha to base64 encoding string.
// mimeType is one of "audio/wav" "image/png".
func WriteToBase64Encoding(cap base64Captcha.CaptchaInterface) string {
	return base64Captcha.CaptchaWriteToBase64Encoding(cap)
}

// NewConfig 根据参数生成验证码参数
func NewConfig(captchaType string, captchaLen, height, width int) interface{} {
	var conf interface{}
	switch strings.ToLower(captchaType) {
	case AudioCaptcha:
		conf = base64Captcha.ConfigAudio{
			CaptchaLen: captchaLen,
			Language:   "zh",
		}
	case CharacterCaptcha:
		conf = base64Captcha.ConfigCharacter{
			Height: height,
			Width:  width,
			//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
			Mode:               base64Captcha.CaptchaModeNumber,
			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
			IsShowHollowLine:   false,
			IsShowNoiseDot:     false,
			IsShowNoiseText:    false,
			IsShowSlimeLine:    false,
			IsShowSineLine:     false,
			CaptchaLen:         captchaLen,
		}
	case DigitCaptcha:
		fallthrough // 继续执行下一个分支
	default:
		conf = base64Captcha.ConfigDigit{
			Height:     height,
			Width:      width,
			MaxSkew:    0.7,
			DotCount:   80,
			CaptchaLen: captchaLen,
		}
	}
	return conf
}
