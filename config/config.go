package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Environment string         `mapstructure:"environment"`
	App         appConfig      `mapstructure:"application"`
	Database    databaseConfig `mapstructure:"database"`
}

type appConfig struct {
	Port int32 `mapstructure:"port"`
}

type databaseConfig struct {
	Name string `mapstructure:"name"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
}

func New() (cfg *Config) {
	cfg = &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("fatal error decoding config file: %w", err))
	}

	return
}
