package config

type OTA struct {
	BuildHost     string `mapstructure:"build-host" json:"buildHost" yaml:"build-host"` // 构建服务器地址
	BuildUser     string `mapstructure:"build-user" json:"buildUser" yaml:"build-user"` // 构建服务器用户名
	BuildPassword string `mapstructure:"build-pass" json:"buildPass" yaml:"build-pass"` // 构建服务器密码

	BuildPath string `mapstructure:"build-path" json:"buildPath" yaml:"build-path"` // 构建服务器构建路径
}
