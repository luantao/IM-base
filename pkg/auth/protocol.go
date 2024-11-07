package auth

type Config struct {
	Endpoint    string // 必填
	NamespaceID string // 必填
	AccessKey   string // 必填
	SecretKey   string // 必填
	GroupName   string // 必填
	ConfigType  string // 必填
}

func (c *Config) validate() bool {
	// nacos config 校验
	if c.Endpoint == "" || c.NamespaceID == "" || c.AccessKey == "" || c.SecretKey == "" {
		return false
	}

	// GroupName 校验
	if c.GroupName == "" {
		return false
	}

	if c.ConfigType != "yaml" {
		return false
	}

	return true
}

type GrantRequest struct {
	AppName string `json:"app_name"` // 应用名称
	OpTime  int64
}

type GrantResponse struct {
	AppID     string
	AppSecret string
}

type AuthenticationRequest struct {
	AppID string
	Token string
}

// AppAuthInfo 应用项
type AppAuthInfo struct {
	AppID       string `json:"app_id" yaml:"app_id"`             // 应用 ID
	AppSecret   string `json:"app_secret" yaml:"app_secret"`     // 应用密钥
	AppName     string `json:"app_name" yaml:"app_name"`         // 应用名称
	ServiceName string `json:"service_name" yaml:"service_name"` // 下游服务名称
	CreatedAt   int64  `json:"created_at" yaml:"created_at"`     // 生成时间
	UpdatedAt   int64  `json:"updated_at" yaml:"updated_at"`     // 更新时间
}

type ParseAuthorizationResponse struct {
	AppID string
	Token string
}

type AppAuthInfos struct {
	Data []AppAuthInfo `yaml:"data"`
}
