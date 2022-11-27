package config

type Signature struct {
	Secret     string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expiration int    `mapstructure:"expiration" json:"expiration" yaml:"expiration"`
	App        string `mapstructure:"app" json:"app" yaml:"app"`
	Version    string `mapstructure:"version" json:"version" yaml:"version"`
	TurnOn     bool   `mapstructure:"turn-on" json:"turn-on" yaml:"turn-on"`
}
