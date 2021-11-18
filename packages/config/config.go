package config

import (
	"github.com/spf13/viper"
	"strings"
)

type (
	App struct {
		Name    string `toml:"name"`
		Mode    string `toml:"mode"`
		Version string `toml:"version"`
	}
	Logger struct {
		Level string `toml:"level"`
	}
	HTTP struct {
		Address string `toml:"address"`
	}
	Config struct {
		App    `toml:"app"`
		Logger `toml:"logger"`
		HTTP   `toml:"http"`
	}
)

var config = Config{}

func init() {
	var err error
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("toml")

	v.AddConfigPath("./temp")
	v.AddConfigPath("./temp/config")
	v.AddConfigPath(".")
	v.AddConfigPath("./conf")
	v.AddConfigPath("./config")

	if err = v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.SetDefault("app.name", "api-server")
	v.SetDefault("app.mode", "debug")
	v.SetDefault("app.version", "1.0.0")
	v.SetDefault("logger.level", "debug")
	v.SetDefault("http.address", ":2000")

	if err = v.Unmarshal(&config); err != nil {
		panic(err)
	}

	config.Logger.Level = strings.ToUpper(config.Logger.Level)
}

func Get() Config {
	return config
}
