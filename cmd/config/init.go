package config

import (
	"os"

	"github.com/spf13/viper"
)

type ProgLang string

const (
	Go         ProgLang = "go"
	Javascript ProgLang = "javascript"
)

type Editor string

const (
	VSCode Editor = "code"
	NeoVim Editor = "nvim"
)

type Config struct {
	Editor    Editor   `mapstructure:"EDITOR"`
	ProgLang  ProgLang `mapstructure:"PROG_LANG"`
	TasksPath string   `mapstructure:"TASKS_PATH"`
}

func Create() (config *Config, err error) {
	viper.AddConfigPath(getEnvLocation())
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func getEnvLocation() string {
	envLocation := os.Getenv("ENV_LOCATION")
	if envLocation != "" {
		return envLocation
	}
	envLocation, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return envLocation
}
