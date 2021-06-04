package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	SelectForeground string `mapstructure:"select_foreground"`
	SelectBackground string `mapstructure:"select_background"`
	Theme            string `mapstructure:"theme"`
}

// Config is a global struct holding the configuration
var Config config

// LoadConfig finds and unmarshals configuration
func LoadConfig() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$HOME/.config/make-tui")
	setDefaults()
	err = viper.ReadInConfig()
	if err != nil {
		// Do not panic if config is not found.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = nil
		} else {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
	}
	viper.Unmarshal(&Config)
	return
}

func setDefaults() {
	viper.SetDefault("select_foreground", "white")
	viper.SetDefault("select_background", "black")
	viper.SetDefault("theme", "vim")
}
