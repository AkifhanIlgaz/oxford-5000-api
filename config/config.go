package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI string `mapstructure:"MONGO_URI"`
	Port     string `mapstructure:"PORT"`
	Mode     string
}

func (c *Config) Validate() error {
	if c.MongoURI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}

	return nil
}

func Load() (Config, error) {
	var config Config

	flag.StringVar(&config.Mode, "mode", "dev", "App mode: dev (default), prod")
	flag.Parse()

	configFile := envFile(config.Mode)

	viper.SetConfigFile(configFile)

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

func envFile(mode string) string {
	switch mode {
	case devMode, prodMode:
		return mode + ".env"
	default:
		return ""
	}
}
