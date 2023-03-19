package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig(config_file string) (*viper.Viper, error) {
	// Créer une nouvelle instance de Viper
	viper := viper.New()

	// Définir le nom du fichier de configuration à charger
	viper.SetConfigFile(config_file)

	// Lire le fichier de configuration
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err)
	}

	return viper, nil
}
