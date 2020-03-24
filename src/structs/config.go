package structs

type SystemConfiguration struct {
	// 服务器配置
	Server ScServer `yaml:"server" json:"server"`
	// 调试模式
	DebugMode bool `yaml:"debug" json:"debug"`
}

type ScServer struct {
	// 监听IP
	Listen string `yaml:"listen" json:"listen"`
	// 监听端口
	Port int `yaml:"port" json:"port,int"`
	// HTML模板
	Templates string `yaml:"templates" json:"templates"`
}
