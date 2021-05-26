package apiserver

import "github.com/Beknur1003/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` //Означает адрес на котором мы запускаем веб-сервер
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config { //NewConfig будет отдавать какой-то уже инициализированный конфиг с дефолтными параметрами
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
