package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// KafkaConfig holds Kafka-related settings.
type KafkaConfig struct {
	Broker string `mapstructure:"broker"`
}

type Config struct {
	Port         string        `mapstructure:"port"`
	Env          string        `mapstructure:"env"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	Kafka        KafkaConfig   `mapstructure:"kafka"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	// Get environment mode from ENV variable (default: "development")
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Set the config file based on environment
	configFile := env + ".yaml"
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // Add current directory

	// Read YAML file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: No config file found: %v", err)
	}

	err = viper.Unmarshal(&config)
	return config, err
}
