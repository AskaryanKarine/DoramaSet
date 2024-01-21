package logger

import (
	"DoramaSet/internal/config"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
	"io"
	"os"
)

type Logger struct {
	*logrus.Entry
	File *os.File
}

func Init(cfg *config.Config) (*Logger, error) {
	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(cfg.Logger.FileName, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	log := logrus.Logger{
		Out: io.Writer(f),
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
		Level: level,
	}
	logrus.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	)))
	l := Logger{
		logrus.NewEntry(&log),
		f,
	}

	return &l, nil
}
