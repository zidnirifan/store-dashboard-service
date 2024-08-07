package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

var config Config

func init() {
	viper := viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config = Config{
		Port: viper.GetInt("PORT"),
	}
}

func GetConfig() *Config {
	return &config
}
