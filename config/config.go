package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string        `mapstructure:"PORT"`
	Env          string        `mapstructure:"ENV"`
	ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("PORT", "8081")
	viper.SetDefault("ENV", "development")
	viper.SetDefault("READ_TIMEOUT", "5s")
	viper.SetDefault("WRITE_TIMEOUT", "10s")

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("No config file found: %v", err)
	}
	err = viper.Unmarshal(&config)
	return config, err
}
