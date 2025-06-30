package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI string `mapstructure:"MONGO_URI"`
	Port     string `mapstructure:"PORT"`

	AccessTokenPrivateKey string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiry     int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`

	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiry     int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`

	Mode string

	RedisConnectionString string `mapstructure:"REDIS_CONNECTION_STRING"`
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
