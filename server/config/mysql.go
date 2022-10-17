package config

import "fmt"

type Mysql struct {
	IP   string `mapstructure:"ip" json:"ip" yaml:"ip"` // server ip address
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// Config       string `mapstructure:"config" json:"config" yaml:"config"`
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // db username
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // db pw
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"` // turn on global gorm log
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`    // use zap to log
	CharSet      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	ParseTime    string `mapstructure:"parseTime" json:"parseTime" yaml:"parseTime"` // to parse time based on locale
	Loc          string `mapstructure:"loc" json:"loc" yaml:"loc"`                   //pass in the locale that used to parse time
}

// data source name
func (m *Mysql) Dsn() string {
	// more params see https://github.com/go-sql-driver/mysql#parameters
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", m.Username, m.Password, m.IP, m.Port, m.DbName, m.CharSet, m.ParseTime, m.Loc)
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
