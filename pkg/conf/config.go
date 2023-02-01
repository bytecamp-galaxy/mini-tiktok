package conf

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	configPathKey = "config"
)

var (
	V *viper.Viper
)

func Init() *viper.Viper {
	if V != nil {
		return V
	}

	v := viper.New()

	pflag.String(configPathKey, "configs/global.yaml", "Path to config file")
	pflag.Parse()
	err := v.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}

	path := v.GetString(configPathKey)
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	V = v
	return V
}
