package config

type Session struct {
	Key                     string `mapstructure:"key" json:"key" yaml:"key"`
	CookieName              string `mapstructure:"cookie-name" json:"cookie-name" yaml:"cookie-name"`
	ExpireTime              int    `mapstructure:"expire-time" json:"expire-time" yaml:"expire-time"`
	RefreshBeforeExpireTime int64  `mapstructure:"refresh-before-expire-time" json:"refresh-before-expire-time" yaml:"refresh-before-expire-time"`
	HttpOnly                bool   `mapstructure:"http-only" json:"http-only" yaml:"http-only"`
	Secure                  bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
}
