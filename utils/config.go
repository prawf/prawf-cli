package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ConfigName = "prawf"
	ConfigType = "json"
)

type PrawfConfig struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

func GetConf(configPath string, configName string) *PrawfConfig {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(ConfigType)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &PrawfConfig{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
