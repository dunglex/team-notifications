package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	WebhookURL          string `mapstructure:"WEBHOOK_URL"`
	HealthCheckURL      string `mapstructure:"HEALTH_CHECK_URL"`
	HealthCheckInterval int    `mapstructure:"HEALTH_CHECK_INTERVAL_SECONDS"`
	Port                string `mapstructure:"PORT"`
}

func LoadConfig() (Config, error) {
	// if .env exists, read from it
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
