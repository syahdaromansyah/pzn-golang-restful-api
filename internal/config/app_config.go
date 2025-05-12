package config

import (
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"
)

type (
	Server struct {
		Host   string
		Port   int
		ApiKey string
	}

	Database struct {
		Username        string
		Password        string
		Host            string
		Port            int
		DBName          string
		MinConns        int
		MaxConns        int
		MaxConnLifeTime time.Duration
		MaxConnIdleTime time.Duration
	}

	Log struct {
		Level     int
		Formatter string
		Output    string
		FilePath  string
	}

	Test struct {
		Timeout time.Duration
	}

	AppConfig struct {
		Server   *Server
		Database *Database
		Log      *Log
		Test     *Test
	}
)

type AppConfigPaths []string

var (
	syncOnce  sync.Once
	appConfig *AppConfig
)

func NewAppConfig(configPaths AppConfigPaths) *AppConfig {
	syncOnce.Do(func() {
		vp := viper.New()

		vp.SetConfigName("config")
		vp.SetConfigType("yaml")

		for _, configPath := range configPaths {
			vp.AddConfigPath(configPath)
		}

		err := vp.ReadInConfig()
		helper.LogStdPanicIfError(err)

		err = vp.Unmarshal(&appConfig)
		helper.LogStdPanicIfError(err)
	})

	return appConfig
}
