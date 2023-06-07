package MakeConfiger

import (
	"flag"
	"github.com/GoSeoTaxi/ParsingAcatOnline/internal/initApp"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Debug      bool `env:"BONUS_APP_SERVER_DEBUG"`
	URLReq     chan string
	DataOUTReq chan []byte
	ListUrl    map[int]string
	Exit       bool
}

// InitConfig initialises config, first from flags, then from env, so that env overwrites flags
func InitConfig() (*Config, error) {

	var cfg Config

	cfg.ListUrl = make(map[int]string)

	initApp.LoadTOmapTXT(cfg.ListUrl)

	flag.BoolVar(&cfg.Debug, "debug", true, "key for hash function")
	flag.Parse()

	err := env.Parse(&cfg)

	cfg.URLReq = make(chan string)
	cfg.DataOUTReq = make(chan []byte)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
