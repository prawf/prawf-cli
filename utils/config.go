package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	ConfigName = "prawf"
	ConfigType = "json"
)

type PrawfConfig struct {
	Tests       map[string]Test `mapstructure:"tests"`
	CurrentTest string          `mapstructure:"current-test"`
}

type Test struct {
	URL     string            `mapstructure:"url"`
	Methods map[string]Method `mapstructure:"methods"`
}

type Method struct {
	Path   string                 `mapstructure:"path"`
	Method string                 `mapstructure:"method"`
	Query  map[string]interface{} `mapstructure:"query,omitempty"`
	Body   map[string]interface{} `mapstructure:"body,omitempty"`
}

func CreateConfigFile(configPath string, configName string) (string, error) {
	fileName := configName + "." + ConfigType
	filePath := filepath.Join(configPath, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return "", err
		}
		err = ioutil.WriteFile(filePath, []byte("{}"), 0644)
		if err != nil {
			return "", err
		}
	}
	return filePath, nil
}

func GetPrawfConfig(v *viper.Viper) (*PrawfConfig, error) {
	conf := &PrawfConfig{}
	err := v.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func AddTestsToConfig(testName string, test Test, filePath string) error {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	conf, err := GetPrawfConfig(viper.GetViper())

	if err != nil {
		return err
	}

	if conf.Tests == nil {
		conf.Tests = map[string]Test{}
	}

	_, exists := conf.Tests[testName]

	if exists {
		return errors.New("error adding new test: A test with the given name already exists")
	}

	conf.Tests[testName] = test

	viper.Set("tests", conf.Tests)
	viper.Set("current-test", testName)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
