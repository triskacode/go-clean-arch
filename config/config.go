package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Environment string    `mapstructure:"environment"`
	App         appConfig `mapstructure:"application"`
}

type appConfig struct {
	Port int32 `mapstructure:"port"`
}

func New() (cfg *config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	cfg = &config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("fatal error decoding config file: %w", err))
	}

	return
}
