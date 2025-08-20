package config

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var cfg *Config

func InitConfig(ctx context.Context) {
	data, err := g.Cfg().Data(ctx)
	if err != nil {
		panic(err)
	}
	err = gconv.Scan(data, &cfg)
	if err != nil {
		panic(err)
	}
	customConfig()
}

func GetConfig() Config {
	if cfg == nil {
		panic("init config first")
	}

	return *cfg
}

func customConfig() {}
