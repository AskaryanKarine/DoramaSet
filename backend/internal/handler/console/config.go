package console

import (
	"DoramaSet/internal/config"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var errorConfigRead = errors.New("can't read config file")

var configPath = flag.String("config", "./configs/config.yml", "config file path")

func initConfig() (*config.Config, error) {
	var cfg config.Config

	flag.Parse()
	viper.SetConfigFile(*configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("readInConfig: %w", errorConfigRead)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return &cfg, nil
}
