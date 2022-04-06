package config

import (
	"github.com/kunaltaitkar/golang-template-project/model"

	"github.com/spf13/viper"
)

func Load(path string) (config model.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
