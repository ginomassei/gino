package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load(configKey string, T interface{}) {
	// Unmarshal config
	err := viper.UnmarshalKey(configKey, &T)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
