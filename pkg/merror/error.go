package merror

import (
	"fmt"

	"go.uber.org/zap"
)

// Level 错误等级
type Level int8

// 错误等级
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
)

// 空的Error
var Nil Error = Error{}

// 错误码对应的错误等级
var levelMap map[Code]Level = map[Code]Level{
	OK:              InfoLevel,
	ServiceError:    ErrorLevel,
	ParamMissing:    ErrorLevel,
	ParamInvalid:    ErrorLevel,
	SignatureFailed: ErrorLevel,
	TokenInvalid:    WarnLevel,
	Forbidden:       WarnLevel,
	RequestLimit:    ErrorLevel,
	BadRequest:      ErrorLevel,
	InternalError:   ErrorLevel,
}

// Error 错误处理
type Error struct {
	code        Code        // 公共错误码
	msg         string      // 公共错误信息
	subCode     int         // 业务错误码
	subMsg      string      // 业务错误信息
	extraFields []zap.Field // 附加的字段信息（用于日志输出）
}

// New 创建新的错误对象
func New(code Code, subCode int, subMsg string) Error {
	return Error{
		code:    code,
		msg:     codeMsgMap[code],
		subCode: subCode,
		subMsg:  subMsg,
	}
}

// Newf 带有参数（类似Sprintf 输出参数）的错误对象
func Newf(code Code, subCode int, subMsg string, args ...interface{}) Error {
	return New(code, subCode, fmt.Sprintf(subMsg, args...))
}

// Code 错误码
func (e Error) Code() Code {
	return e.code
}

// Msg 公共错误信息
func (e Error) Msg() string {
	return e.msg
}

// SubCode 业务错误码
func (e Error) SubCode() int {
	return e.subCode
}

// SubMsg 业务错误信息
func (e Error) SubMsg() string {
	return e.subMsg
}

// PrependSubMsg 在前面添加错误信息，生成新的Error（格式 msg:err.SubMsg）
func (e Error) PrependSubMsg(msg string) Error {
	e.subMsg = msg + ":" + e.subMsg
	return e
}

// PrependSubMsgf 参数化的PrependSubMsg
func (e Error) PrependSubMsgf(msg string, args ...interface{}) Error {
	msg = fmt.Sprintf(msg, args...)
	return e.PrependSubMsg(msg)
}

// AppendSubMsg 添加错误信息,生成新的Error（格式 err.SubMsg:msg）
func (e Error) AppendSubMsg(msg string) Error {
	e.subMsg = e.SubMsg() + ":" + msg
	return e
}

// AppendSubMsgf 参数化的AppendSubMsg
func (e Error) AppendSubMsgf(msg string, args ...interface{}) Error {
	msg = fmt.Sprintf(msg, args...)
	return e.AppendSubMsg(msg)
}

// ResetCode 重置一级错误码
func (e Error) ResetCode(code Code) Error {
	e.code = code
	e.msg = codeMsgMap[code]
	return e
}

// ResetSubCode 重置业务错误码
func (e Error) ResetSubCode(subCode int) Error {
	e.subCode = subCode
	return e
}

// ResetSubMsg 重新设置Submsg错误信息生成一个新的Error
func (e Error) ResetSubMsg(msg string) Error {
	e.subMsg = msg
	return e
}

// ResetSubMsgf 参数化的ResetSubMsg
func (e Error) ResetSubMsgf(msg string, args ...interface{}) Error {
	e.subMsg = fmt.Sprintf(msg, args...)
	return e
}

// AppendExtraField 添加附加的字段信息（用于日志输出）
func (e Error) AppendExtraField(fields ...zap.Field) Error {
	e.extraFields = append(e.extraFields, fields...)
	return e
}

// ExtraField 获取附加的字段信息
func (e Error) ExtraFields() []zap.Field {
	return e.extraFields
}

// Error 错误信息
func (e Error) Error() string {
	if e.IsNil() {
		return ""
	}
	return fmt.Sprintf("%#v", e)
}

// GetError 获取error类型的错误
func (e Error) GetError() error {
	if e.IsNil() {
		return nil
	}
	return e
}

// IsNil 是否有错误
func (e Error) IsNil() bool {
	return e.code == 0 || e.code == OK
}

// IsNotNil 是否没有错误
func (e Error) IsNotNil() bool {
	return !e.IsNil()
}

// Level 错误级别
func (e Error) Level() Level {
	level, ok := levelMap[e.code]
	// 默认为Info
	if !ok {
		return InfoLevel
	}
	return level
}
