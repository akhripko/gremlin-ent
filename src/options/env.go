package options

import (
	"github.com/spf13/viper"
)

func ReadEnv() *Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("GREMLIN_ADDR", "localhost:8182")

	return &Config{
		GremlinAddr: viper.GetString("GREMLIN_ADDR"),
	}
}
