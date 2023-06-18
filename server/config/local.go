package config

type Local struct {
	StaticFilePath string `mapstructure:"static-file-path" json:"static-file-path" yaml:"static-file-path"`
}
