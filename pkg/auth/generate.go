package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/luantao/IM-base/pkg/id"
	"github.com/luantao/IM-base/pkg/merror"
	"strings"
	"time"
)

// generateAppID "auth_"+ appName + _ + 16位随机数
func generateAppID(appName string) string {
	return fmt.Sprintf("auth_%s_%s", appName, id.RandomID())
}

// generateAppSecret
// sha1(app_name+service_name+当前时间戳) 后转为大写
func generateAppSecret(appName, serviceName string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%s%d", appName, serviceName, time.Now().Unix())))
	return hex.EncodeToString(h.Sum(nil))
}

func GenAuthorizationInfo(appID, appSecret string) string {
	var buf strings.Builder
	buf.WriteString("app_id=")
	buf.WriteString(appID)
	buf.WriteString("&")
	buf.WriteString("token=")
	buf.WriteString(GenToken(appID, appSecret))
	authorizationStr := buf.String()
	return authorizationStr
}

func GenToken(appID, appSecret string) (token string) {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s%s", appID, appSecret)))
	hStr := hex.EncodeToString(h.Sum(nil))
	token = hStr
	return
}

func ParseAuthorization(authorizationStr string) (result ParseAuthorizationResponse, ygErr merror.Error) {
	temp := strings.Split(authorizationStr, "&")
	if len(temp) != 2 {
		ygErr = AuthorizationErr
		return
	}
	appIDInfo := strings.Split(temp[0], "=")
	tokenInfo := strings.Split(temp[1], "=")
	// 长度不符合
	if len(appIDInfo) != 2 || len(tokenInfo) != 2 {
		if len(temp) != 2 {
			ygErr = AuthorizationErr
			return
		}
	}
	result.AppID = appIDInfo[1]
	result.Token = tokenInfo[1]
	return
}
