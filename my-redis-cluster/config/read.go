package config

import (
	"github.com/spf13/viper"
	"strconv"
)
import log "github.com/sirupsen/logrus"

func ReadConfig() {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath("./config")

	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config := &Config{}
	_ = v.Unmarshal(&config)
	log.Info("我是参数: " + strconv.Itoa(config.AppConfig.Port))
	log.Info("read config yml success!")
	return
}
