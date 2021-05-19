package utils

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

const (
	ConfigName = "prawf"
	ConfigType = "json"
)

type PrawfConfig struct {
	Current string          `json:"current"`
	Tests   map[string]Test `json:"tests"`
}

type Test struct {
	URL     string            `json:"url"`
	Methods map[string]Method `json:"methods"`
}

type Method struct {
	Path   string                 `json:"path"`
	Method string                 `json:"method"`
	Header map[string]interface{} `json:"header,omitempty"`
	Query  map[string]interface{} `json:"query,omitempty"`
	Body   map[string]interface{} `json:"body,omitempty"`
}

func (pc *PrawfConfig) GetTest() (Test, error) {
	if test, e := pc.Tests[pc.Current]; e {
		return test, nil
	}

	return Test{}, errors.New("Current test " + pc.Current + " not defined.")
}

func CreateConfigFile(filePath string) error {

	_, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, []byte("{}"), 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetPrawfConfig(v *viper.Viper) (*PrawfConfig, error) {
	conf := &PrawfConfig{}
	err := v.Unmarshal(&conf)
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
	viper.Set("current", testName)

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
