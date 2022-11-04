package config

import (
	"os"

	"github.com/spf13/viper"
)

type ApiConfig struct {
	Port string
}

type DatabaseConfig struct {
	Port     int
	Host     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	Api      ApiConfig
	Database DatabaseConfig
}

var cfg *Config

func LoadConfig() {
	//Default dev
	viper.SetDefault("env", "dev")

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(os.Getenv("ENV"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(cfg)

	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return cfg
}
