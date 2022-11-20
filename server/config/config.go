package config

type Server struct {
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Session Session `mapstructure:"session" json:"session" yaml:"session"`
	Lock    Lock    `mapstructure:"lock" json:"lock" yaml:"lock"`
	Limiter Limiter `mapstructure:"limiter" json:"limiter" yaml:"limiter"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
