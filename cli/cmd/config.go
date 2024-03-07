package cmd

// 负责启动时候的config读取
type CliConfig struct {
}

// 读取返回config结构
func InitConfig() *CliConfig {
	return &CliConfig{}
}
