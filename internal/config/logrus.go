package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogrus(vp *viper.Viper) *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.Level(vp.GetInt("log.level")))

	if vp.GetString("log.formatter") == "json" {
		logger.SetFormatter(new(logrus.JSONFormatter))
	}

	return logger
}
