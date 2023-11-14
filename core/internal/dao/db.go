package dao

import (
	"fmt"

	"github.com/krau/manyacg/core/internal/common/config"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Database,
	)
	theDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
		
	})
	if err != nil {
		panic(err)
	}
	theDb.AutoMigrate(&entityModel.Artwork{}, &entityModel.Picture{}, &entityModel.Tag{})
	db = theDb
}
