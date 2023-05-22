package config

import "time"

// Config represents the configuration for the URL monitoring tool.
type Config struct {
	MonitoringInterval time.Duration
	SenderEmail        string
	RecipientEmail     string
	EmailService       string
}

// LoadConfig loads the configuration with default values.
func LoadConfig() *Config {
	return &Config{
		MonitoringInterval: 1 * time.Minute,
		SenderEmail:        "aud1sm0@live.com",
		RecipientEmail:     "aud1sm0@live.com",
		EmailService:       "sendinblue", // Replace with your chosen email service
	}
}
