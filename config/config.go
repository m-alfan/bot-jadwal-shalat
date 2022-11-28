package config

import (
	"github.com/joeshaw/envdecode"
)

type Config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN"`
}

func NewConfig() *Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
