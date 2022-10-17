package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// add one layer in order to rewrite its Printf function
type Writer struct {
	logger.Writer
}

// constructor function
func NewWriter(w logger.Writer) *Writer {
	return &Writer{w}
}

// rewrite Printf func of Writer interface
func (w *Writer) Printf(message string, data ...interface{}) {
	logZap := global.CONFIG.Mysql.LogZap
	if logZap {
		global.LOGGER.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

func GormInit() *gorm.DB {
	m := global.CONFIG.Mysql

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false, //  auto configure based on currently MySQL version
	}
	global.LOGGER.Info(fmt.Sprintf("dsn: %v\n", mysqlConfig.DSN))

	db, err := gorm.Open(mysql.New(mysqlConfig), GetGormConfig())
	if err != nil {
		global.LOGGER.Error(err.Error())
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db
}

func GetGormConfig() *gorm.Config {
	config := gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_writer := NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags))

	_loggerConfig := logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	}

	config.Logger = logger.New(_writer, _loggerConfig)
	return &config
}
