package config

type Limiter struct {
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Interval int    `mapstructure:"interval" json:"interval" yaml:"interval"`
	Limit    int    `mapstructure:"limit" json:"limit" yaml:"limit"`
}
