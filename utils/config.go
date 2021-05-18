package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ConfigName = "prawf"
	ConfigType = "json"
)

type config struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

func GetConf(configPath string, configName string) *config {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(ConfigType)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
