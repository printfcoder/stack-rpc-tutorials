package config

// EtcdConfig Etcd 配置
type EtcdConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// defaultEtcdConfig认 etcd 配置
type defaultEtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// GetPort consul 端口
func (c defaultEtcdConfig) GetPort() int {
	return c.Port
}

// GetEnabled consul 激活
func (c defaultEtcdConfig) GetEnabled() bool {
	return c.Enabled
}

// GetHost consul 主机地址
func (c defaultEtcdConfig) GetHost() string {
	return c.Host
}
