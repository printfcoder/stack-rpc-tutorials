package config

// ConsulConfig consul 配置
type ConsulConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// defaultConsulConfig 默认consul 配置
type defaultConsulConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// GetPort consul 端口
func (c defaultConsulConfig) GetPort() int {
	return c.Port
}

// GetEnabled consul 激活
func (c defaultConsulConfig) GetEnabled() bool {
	return c.Enabled
}

// GetHost consul 主机地址
func (c defaultConsulConfig) GetHost() string {
	return c.Host
}
