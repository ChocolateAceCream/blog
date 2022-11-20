package config

type Captcha struct {
	Height     int     `mapstructure:"height" json:"height" yaml:"height"`
	Width      int     `mapstructure:"width" json:"width" yaml:"width"`
	Length     int     `mapstructure:"length" json:"length" yaml:"length"`
	MaxSkew    float64 `mapstructure:"max-skew" json:"max-skew" yaml:"max-skew"`
	DotCount   int     `mapstructure:"dot-count" json:"dot-count" yaml:"dot-count"`
	Expiration int     `mapstructure:"expiration" json:"expiration" yaml:"expiration"`
	Prefix     string  `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}
