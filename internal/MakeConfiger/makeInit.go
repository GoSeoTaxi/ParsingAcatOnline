package MakeConfiger

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Debug bool `env:"BONUS_APP_SERVER_DEBUG"`
}

// InitConfig initialises config, first from flags, then from env, so that env overwrites flags
func InitConfig() (*Config, error) {

	var cfg Config
	flag.BoolVar(&cfg.Debug, "debug", true, "key for hash function")
	flag.Parse()

	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
