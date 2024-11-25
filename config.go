package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
	config.WebhookURL = os.Getenv("WEBHOOK_URL")
	config.HealthCheckURL = os.Getenv("HEALTH_CHECK_URL")
	interval, err := strconv.ParseUint(os.Getenv("HEALTH_CHECK_INTERVAL_SECONDS"), 0, 32)
	if err != nil {
		return fmt.Errorf("invalid HEALTH_CHECK_INTERVAL_SECONDS: %v", err)
	}
	config.HealthCheckInterval = uint(interval)
	config.Port = os.Getenv("PORT")

	// if .env exists, read from it
	if _, err := os.Stat(".env"); err == nil {
		fmt.Println("Reading from .env file")
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
