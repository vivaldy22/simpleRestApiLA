package viper

import (
	"github.com/spf13/viper"
)

func GetEnv(key, defaultValue string) string {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
