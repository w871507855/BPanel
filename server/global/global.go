package global

import (
	"github.com/w871507855/BPanel/server/config"
	"gorm.io/gorm"
)

var (
	CONF config.Server
	DB   *gorm.DB
)
