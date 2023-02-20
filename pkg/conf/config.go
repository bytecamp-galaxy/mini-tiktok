package conf

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

const (
	configPathKey = "config"
	envKey        = "env"
)

const (
	dev  = "dev"
	prod = "prod"
)

var (
	V *viper.Viper
)

func IsDev() bool {
	return Init().GetString("envKey") == dev
}

func IsProd() bool {
	return Init().GetString("envKey") == prod
}

func Init() *viper.Viper {
	if V != nil {
		return V
	}

	v := viper.New()

	// parse config path
	pflag.String(configPathKey, "configs/global.yaml", "Path to config file")
	pflag.Parse()
	err := v.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}

	// read config
	path := v.GetString(configPathKey)
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// get env
	err = v.BindEnv(envKey, "MINI_TIKTOK_ENV")
	if err != nil {
		log.Panic(err)
	}

	V = v
	return V
}
