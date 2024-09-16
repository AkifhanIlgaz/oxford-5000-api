package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI                     string `mapstructure:"MONGO_URI"`
	GoogleApplicationCredentials string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
}

func (c *Config) Validate() error {
	if c.MongoURI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}

	if c.GoogleApplicationCredentials == "" {
		return fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS is required")
	}

	return nil
}

func Load() (Config, error) {
	var config Config

	viper.SetConfigFile("app.env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := config.Validate(); err != nil {
		return config, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}
