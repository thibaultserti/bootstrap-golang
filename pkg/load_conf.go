package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Message string `mapstructure:"MESSAGE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err_aa error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err_aa = viper.ReadInConfig()
	if err_aa != nil {
		return
	}

	err_aa = viper.Unmarshal(&config)
	return
}
