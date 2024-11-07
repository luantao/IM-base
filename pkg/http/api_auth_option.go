package http

// ApiAuthOption 可选参数类型
type ApiAuthOption func(*apiAuthOption)

type apiAuthOption struct {
	appID     string
	appSecret string
}

// applyOpts 应用可选参数
func (op *apiAuthOption) applyOpts(opt ApiAuthOption) {
	opt(op)
}

func ApiAuth(appID, appSecret string) ApiAuthOption {
	return func(option *apiAuthOption) {
		option.appID = appID
		option.appSecret = appSecret
	}
}
