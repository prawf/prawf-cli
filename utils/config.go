package utils

import (
	"errors"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ConfigName = "prawf"
	ConfigType = "json"
)

// PrawfConfig represents the structure of the config file (prawf.json)
type PrawfConfig struct {
	Current string          `json:"current"`
	Tests   map[string]Test `json:"tests"`
}

type Test struct {
	URL     string   `json:"url"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name   string                 `json:"name"`
	Path   string                 `json:"path"`
	Method string                 `json:"method"`
	Header map[string]interface{} `json:"header,omitempty"`
	Query  map[string]interface{} `json:"query,omitempty"`
	Body   map[string]interface{} `json:"body,omitempty"`
	Expect Expect                 `json:"expect,omitempty"`
}

type Expect struct {
	Contain map[string]interface{} `json:"contain,omitempty"`
	Keys    []string               `json:"keys,omitempty"`
	Equal   map[string]interface{} `json:"equal,omitempty"`
}

// (e *Expect) Print prints the expected response in the specifed format
// TODO: Too complicated and WET code. Make it DRY
func (e *Expect) Print(testResult string, outputFmt string) {
	if e.Equal != nil {
		el := log.New().WithField("type", "equal")
		if outputFmt == "json" {
			el.Logger.SetFormatter(&log.JSONFormatter{})
		}
		if testResult == "pass" {
			el.Info(ToJSONToString(e.Equal))
		} else {
			el.Error(ToJSONToString(e.Equal))
		}
	}
	if e.Contain != nil {
		cl := log.New().WithField("type", "contain")
		if outputFmt == "json" {
			cl.Logger.SetFormatter(&log.JSONFormatter{})
		}
		if testResult == "pass" {
			cl.Info(ToJSONToString(e.Contain))
		} else {
			cl.Error(ToJSONToString(e.Contain))
		}
	}
	if e.Keys != nil {
		kl := log.New().WithField("type", "keys")
		if outputFmt == "json" {
			kl.Logger.SetFormatter(&log.JSONFormatter{})
		}
		if testResult == "pass" {
			kl.Info(e.Keys)
		} else {
			kl.Error(e.Keys)
		}
	}
}

// GetTest returns the configuration of the test mentioned in current
func (pc *PrawfConfig) GetTest() (Test, error) {
	if test, e := pc.Tests[pc.Current]; e {
		return test, nil
	}

	return Test{}, errors.New("current test " + pc.Current + " not defined")
}

// CreateConfigFile will create a new config file in the specified path
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

// GetPrawfConfig returns the config file
func GetPrawfConfig(v *viper.Viper) (*PrawfConfig, error) {
	conf := &PrawfConfig{}
	err := v.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// AddTestsToConfig will add a new test to the config file
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
