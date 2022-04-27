package v1

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (db *gorm.DB, err error) {

	db, err = gorm.Open(sqlite.Open("chk.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return
	}

	if err = db.AutoMigrate(&CSV{}); err != nil {
		return
	}

	return
}
