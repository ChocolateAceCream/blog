package config

type Minio struct {
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Bucket   string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region   string `mapstructure:"region" json:"region" yaml:"region"`
	Https    bool   `mapstructure:"https" json:"https" yaml:"https"`
}
