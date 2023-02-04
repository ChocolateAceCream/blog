package config

type Init struct {
	AdminPassword string `mapstructure:"adminPassword" json:"adminPassword" yaml:"adminPassword"`
	AdminEmail    string `mapstructure:"adminEmail" json:"adminEmail" yaml:"adminEmail"`
	GuestPassword string `mapstructure:"guestPassword" json:"guestPassword" yaml:"guestPassword"`
	GuestEmail    string `mapstructure:"guestEmail" json:"guestEmail" yaml:"guestEmail"`
}
