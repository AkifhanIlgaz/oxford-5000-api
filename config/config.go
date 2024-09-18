package config

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
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

	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	fullPath := filepath.Join(home, c.GoogleApplicationCredentials)

	fmt.Println(fullPath)

	c.GoogleApplicationCredentials = fullPath

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
