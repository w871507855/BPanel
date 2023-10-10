package initialize

import (
	"fmt"

	"github.com/w871507855/BPanel/server/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONF.DB.Username,
		global.CONF.DB.Password,
		global.CONF.DB.Host,
		global.CONF.DB.Port,
		global.CONF.DB.Dataname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate()
	if err != nil {
		panic(err.Error())
	}
	return db
}
