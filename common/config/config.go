package config

// 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
}

// 服务器启动端口配置模型
type Server struct {
	Port string `mapstructure:"port"`
}

// 数据库连接配置模型 数据源配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}
