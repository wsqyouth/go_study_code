package conf

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	RunMode string
	Addr    string
}

var Conf *Config

func Init(configPath string) error {
	err := initConfig(configPath)
	if err != nil {
		return err
	}
	return nil
}

func initConfig(configPath string) error {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config.local")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("napp")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	err := viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}

	watchConfig()

	return nil
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
