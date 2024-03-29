package db

import (
	"os"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// create table, used in initializer
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		gormadapter.CasbinRule{},
		dbTable.User{},
		dbTable.Role{},
		dbTable.Menu{},
		dbTable.Endpoint{},
		dbTable.Article{},
		dbTable.Comment{},
		dbTable.CommentLiker{},
		dbTable.Reply{},
		dbTable.ReplyLiker{},
		dbTable.Notification{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register table success")
}
