package merror

type Code uint32

const (
	// 标准错误码
	OK              Code = 10000 // 成功 (info)
	ServiceError    Code = 10001 // 业务处理异常 (error)
	ParamMissing    Code = 10002 // 参数缺失 (error)
	ParamInvalid    Code = 10003 // 非法参数 (error)
	SignatureFailed Code = 10004 // 签名失败 (error)
	TokenInvalid    Code = 10005 // 令牌过期 (warn)
	Forbidden       Code = 10006 // 业务处理拒绝 (warn)
	RequestLimit    Code = 10007 // 调用超限 (error)
	BadRequest      Code = 10008 // 不合法调用 (error)
	InternalError   Code = 10009 // 系统异常 (error)
)

var codeMsgMap map[Code]string = map[Code]string{
	OK:              "成功",
	ServiceError:    "业务处理异常",
	ParamMissing:    "参数缺失",
	ParamInvalid:    "非法参数",
	SignatureFailed: "签名失败",
	TokenInvalid:    "令牌过期",
	Forbidden:       "业务处理拒绝",
	RequestLimit:    "调用超限",
	BadRequest:      "不合法调用",
	InternalError:   "系统异常",
}
