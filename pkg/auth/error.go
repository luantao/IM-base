package auth

import (
	"github.com/luantao/IM-base/pkg/merror"
)

const (
	JsonUnmarshalErrCode   = 1 // json反序列化失败
	JsonMarshalErrCode     = 2 // json序列化失败
	GetConfigErrCode       = 3 // 操作nacos读失败
	TokenValidateErrCode   = 4 // token 异常
	PublishConfigErrCode   = 5 // 操作nacos写失败
	ParamEmptyErrCode      = 6 // 认证信息为空
	ConfigErrCode          = 7 // config赋值有误请检查
	NewConfigClientErrCode = 8 // nacos初始化失败
	AuthorizationErrCode   = 9 // authorizationStr 错误
)

var (
	GetConfigErr       = merror.New(merror.ServiceError, GetConfigErrCode, "操作nacos读失败")
	JsonUnmarshalErr   = merror.New(merror.ServiceError, JsonUnmarshalErrCode, "json反序列化失败")
	JsonMarshalErr     = merror.New(merror.ServiceError, JsonMarshalErrCode, "json序列化失败")
	TokenValidateErr   = merror.New(merror.TokenInvalid, TokenValidateErrCode, "token校验失败")
	PublishConfigErr   = merror.New(merror.ServiceError, PublishConfigErrCode, "操作nacos写失败")
	ParamEmptyErr      = merror.New(merror.ParamMissing, ParamEmptyErrCode, "认证信息为空")
	ConfigErr          = merror.New(merror.ParamMissing, ConfigErrCode, "config赋值有误请检查")
	NewConfigClientErr = merror.New(merror.ServiceError, NewConfigClientErrCode, "nacos初始化失败")
	AuthorizationErr   = merror.New(merror.ServiceError, AuthorizationErrCode, "authorization信息解析错误")
)
