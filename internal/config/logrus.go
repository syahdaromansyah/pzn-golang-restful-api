package config

import (
	"github.com/sirupsen/logrus"
)

func NewLogrus(appConfig *AppConfig) *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.Level(appConfig.Log.Level))

	if appConfig.Log.Formatter == "json" {
		logger.SetFormatter(new(logrus.JSONFormatter))
	}

	return logger
}
