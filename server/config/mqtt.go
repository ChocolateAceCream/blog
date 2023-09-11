package config

type Mqtt struct {
	Host    string `mapstructure:"host" json:"host" yaml:"host"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	TlsPort string `mapstructure:"tls-Port" json:"tls-Port" yaml:"tls-Port"`
}
