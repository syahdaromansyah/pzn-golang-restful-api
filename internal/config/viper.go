package config

import (
	"sync"

	"github.com/syahdaromansyah/pzn-golang-restful-api/internal/helper"

	"github.com/spf13/viper"
)

type ViperConfigPaths []string

var (
	once sync.Once
	vp   *viper.Viper
)

func NewViper(viperConfigPaths ViperConfigPaths) *viper.Viper {
	once.Do(func() {
		vp = viper.New()

		vp.SetConfigName("config")
		vp.SetConfigType("yaml")

		for _, configPath := range viperConfigPaths {
			vp.AddConfigPath(configPath)
		}

		err := vp.ReadInConfig()
		helper.LogStdPanicIfError(err)
	})

	return vp
}
