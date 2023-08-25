package btchash160

import (
	"btchash160/pkg/logger"

	"github.com/spf13/viper"
)

type config struct {
	Hash160 string `yaml:"hash160"`
	Min     string `yaml:"min"`
	Max     string `yaml:"max"`
}

var conf *config = &config{}

func init() {
	loadConfig()
}

func loadConfig() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Log(err, logger.PanicLevel)
	}
	viper.Unmarshal(conf)
}
