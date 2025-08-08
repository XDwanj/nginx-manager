package config

// DefaultConfig 默认配置结构体
type DefaultConfig struct {
	ServerItems        []string `json:"server_items"`         // 服务器块的默认配置项
	LocationFirstItems []string `json:"location_first_items"` // location 块的第一个配置项
}

// ServiceConfig 服务配置结构体
type ServiceConfig struct {
	ServerName string     `json:"server_name"` // 服务器名称
	Locations  []Location `json:"locations"`   // location 配置列表
}

// Location location 配置结构体
type Location struct {
	Location  string   `json:"location"`   // location 路径
	ProxyPass string   `json:"proxy_pass"` // 代理地址
	Items     []string `json:"items"`      // 其他配置项
}
