package db

import (
	"os"

	"github.com/ChocolateAceCream/blog/db/model"
	"github.com/ChocolateAceCream/blog/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// create table, used in initializer
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		model.User{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
