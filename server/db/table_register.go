package db

import (
	"os"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// create table, used in initializer
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// user model
		dbTable.User{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
