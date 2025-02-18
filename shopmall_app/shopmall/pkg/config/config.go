package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:"server"`
}

type Server struct {
	HTTP HTTP `mapstructure:"http"`
}

type HTTP struct {
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
