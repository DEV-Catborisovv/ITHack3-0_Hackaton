package configs

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Telegram struct {
		BotApiToken string
		Debug       bool
	}
	APIClient struct {
		SecretKey  string
		MerchantId string
	}
}

func NewConfig() *Config {
	s := Config{}
	dat, err := os.ReadFile("config.toml")
	if err != nil {
		return &Config{}
	}

	err = toml.Unmarshal(dat, &s)
	if err != nil {
		return &Config{}
	}
	return &s
}
