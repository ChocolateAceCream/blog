package config

type Email struct {
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	Host       string `mapstructure:"host" json:"host" yaml:"host"`
	Username   string `mapstructure:"username" json:"username" yaml:"username"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
	Expiration int    `mapstructure:"expiration" json:"expiration" yaml:"expiration"`
	Prefix     string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}
