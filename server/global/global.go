package global

import (
	"github.com/ChocolateAceCream/blog/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// global variables
var (
	CONFIG config.Server
	VIPER  *viper.Viper
	LOGGER *zap.Logger
)
