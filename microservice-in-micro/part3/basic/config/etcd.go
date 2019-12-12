package config

// EtcdConfig etcd 配置
type EtcdConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// defaultEtcdConfig 默认etcd 配置
type defaultEtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// GetPort etcd 端口
func (c defaultEtcdConfig) GetPort() int {
	return c.Port
}

// GetEnabled etcd 激活
func (c defaultEtcdConfig) GetEnabled() bool {
	return c.Enabled
}

// GetHost etcd 主机地址
func (c defaultEtcdConfig) GetHost() string {
	return c.Host
}
