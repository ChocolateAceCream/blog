package config

type Redis struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
