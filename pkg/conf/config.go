package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	V *viper.Viper
}

func Init() Config {
	v := viper.New()
	config := Config{V: v}

	v.SetConfigName("global")
	v.SetConfigType("yaml")
	v.AddConfigPath("conf")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return config
}
