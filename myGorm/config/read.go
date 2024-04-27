package config

import (
	"github.com/spf13/viper"
)

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
}
