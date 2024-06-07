package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MONGO_URI string
	DB_NAME   string
}

func LoadConfig() (Config, error) {
	// Create a new Viper instance.
	v := viper.New()

	// Set the configuration file name, path, and environment variable settings.
	v.SetConfigFile(".env")

	// Read the configuration file.
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	// Unmarshal the configuration into the Config struct.
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return Config{}, err
	}

	return config, nil
}
