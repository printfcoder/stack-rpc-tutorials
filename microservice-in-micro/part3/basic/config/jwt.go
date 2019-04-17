package config

// jwtConfig jwt 配置 接口
type JwtConfig interface {
	GetSecretKey() string
}

// defaultJwtConfig jwt 配置
type defaultJwtConfig struct {
	SecretKey string `json:"secretKey"`
}

// GetSecretKey jwt 密钥
func (m defaultJwtConfig) GetSecretKey() string {
	return m.SecretKey
}
