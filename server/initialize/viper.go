package initialize

import (
	"log"

	"github.com/spf13/viper"
	"github.com/w871507855/BPanel/server/global"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("配置文件加载失败, %s", err)
	}
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		log.Fatalf("配置文件格式化有问题, %s", err)
	}
}
