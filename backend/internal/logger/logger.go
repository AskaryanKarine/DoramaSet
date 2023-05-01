package logger

import (
	"DoramaSet/internal/config"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
)

func Init(cfg *config.Config) (*logrus.Logger, *os.File, error) {
	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		return nil, nil, err
	}

	f, err := os.OpenFile(cfg.Logger.FileName, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		_ = f.Close()
		return nil, nil, err
	}
	log := logrus.Logger{
		Out: io.Writer(f),
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
		Level: level,
	}
	return &log, f, nil
}
