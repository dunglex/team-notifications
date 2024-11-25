package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	WebhookURL          string `mapstructure:"WEBHOOK_URL"`
	HealthCheckURL      string `mapstructure:"HEALTH_CHECK_URL"`
	HealthCheckInterval uint   `mapstructure:"HEALTH_CHECK_INTERVAL_SECONDS"`
	Port                string `mapstructure:"PORT"`
}

func (config *AppConfig) LoadConfig() error {
	// read from environment variables
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
	config = &AppConfig{
		WebhookURL:          viper.GetString("WEBHOOK_URL"),
		HealthCheckURL:      viper.GetString("HEALTH_CHECK_URL"),
		HealthCheckInterval: viper.GetUint("HEALTH_CHECK_INTERVAL_SECONDS"),
		Port:                viper.GetString("PORT"),
	}

	// if .env exists, read from it
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			return err
		}
	}

	// manually validate the config
	if config.WebhookURL == "" {
		return fmt.Errorf("WEBHOOK_URL is required")
	}
	if config.HealthCheckURL == "" {
		return fmt.Errorf("HEALTH_CHECK_URL is required")
	}
	if config.HealthCheckInterval == 0 {
		return fmt.Errorf("HEALTH_CHECK_INTERVAL_SECONDS is required")
	}
	if config.Port == "" {
		return fmt.Errorf("PORT is required")
	}

	return nil
}
