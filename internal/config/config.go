package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ApiConfig struct {
	Port int
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

func LoadConfig() *Config {
	//Default dev
	viper.SetDefault("env", "dev")

	viper.AddConfigPath(".")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName(viper.GetString("env"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Panic().Err(err).Msg("Error to parse env vars")
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		log.Panic().Err(err).Msg("Error to decode env vars")
	}

	return cfg
}

func GetConfig() *Config {
	return cfg
}
