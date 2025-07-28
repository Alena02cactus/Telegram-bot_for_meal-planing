package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"TELEGRAM_TOKEN"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
