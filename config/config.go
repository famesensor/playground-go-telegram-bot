package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Telegram Telegram
}

type Telegram struct {
	Token string
}

func NewConfig() *Config {
	viper.SetConfigName("config") // config.yaml, config.json, etc.
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Look for config in current directory
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // Read environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Config file not found, using defaults")
		return nil
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Println("Config Error:", err)
		return nil
	}

	return &config
}
