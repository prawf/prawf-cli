/*
Copyright © 2021 Navendu Pottekkat <navendupottekkat@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/prawf/prawf-cli/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Compare tests with expected response defined in the prawf.json file",
	Long: `Compare tests with expected response defined in the prawf.json file

Uses the test marked as 'current' by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load the config file
		conf, err := utils.GetPrawfConfig(viper.GetViper())

		if err != nil {
			log.Fatal(err)
		}

		test, err := conf.GetTest()

		if err != nil {
			log.Fatal(err)
		}

		log.WithFields(log.Fields{
			"test": conf.Current,
			"url":  test.URL,
		}).Info("Running test.")

		// Perform all the tests mentioned in the config file
		for _, method := range test.Methods {
			fmt.Print("\n")
			TestResponse(test.URL, method.Path, method.Method, method.Header, method.Query, method.Body, method.Name, method.Expect)
		}
	},
}

func TestResponse(
	url string,
	path string,
	method string,
	header map[string]interface{},
	query map[string]interface{},
	body map[string]interface{},
	name string,
	expect utils.Expect) {
	var data interface{}

	// Create a new HTTP request
	req := NewRequest(url, path, method, header, query, body, name)

	// Create a client and make the request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// If the response is HTML, then most likely an error so return that
	if utils.ContentTypeIsHTML(resp) {
		log.WithField("response", resp).Error()
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Log the response
	log.WithFields(log.Fields{
		"status code": resp.Status,
	}).Info("Response received.")

	log.Info(string(b))

	// Get the json data
	err = json.Unmarshal([]byte(b), &data)

	if err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(utils.Expect{}, expect) {
		var testResult string
		if AssertResponse(data, expect) {
			testResult = "pass"
		} else {
			testResult = "fail"
		}

		el := log.New().WithFields(log.Fields{"test": testResult})

		if expect.Equal != nil {
			el = el.WithFields(log.Fields{"equal": "yes"})
		} else {
			el = el.WithFields(log.Fields{"equal": "no"})
		}
		if expect.Contain != nil {
			el = el.WithFields(log.Fields{"contain": "yes"})
		} else {
			el = el.WithFields(log.Fields{"contain": "no"})
		}
		if expect.Keys != nil {
			el = el.WithFields(log.Fields{"keys": "yes"})
		} else {
			el = el.WithFields(log.Fields{"keys": "no"})
		}

		el.Info("Expected response.")
		expect.Print()
	}

}

func AssertResponse(data interface{}, expect utils.Expect) bool {
	// The response could be an interface or an array of interface
	if m, ok := data.([]interface{}); ok {
		if expect.Equal != nil {
			return AssertEqual(m, expect.Equal)
		}
		for _, item := range m {
			if expect.Contain != nil {
				if AssertContain(item.(map[string]interface{}), expect) {
					return true
				}
			}
			if expect.Keys != nil {
				if AssertKeys(item.(map[string]interface{}), expect) {
					return true
				}
			}
		}
		return false
	} else if item, ok := data.(map[string]interface{}); ok {
		if expect.Equal != nil {
			return AssertEqual(item, expect.Equal)
		}
		if expect.Contain != nil {
			return AssertContain(item, expect)
		}
		if expect.Keys != nil {
			return AssertKeys(item, expect)
		}
	} else {
		log.Fatal("Invalid response received.")
	}
	return false
}

func AssertEqual(r interface{}, e interface{}) bool {
	return reflect.DeepEqual(r, e)
}

func AssertContain(item map[string]interface{}, expect utils.Expect) bool {
	for key, value := range expect.Contain {
		if rValue, ok := item[key]; ok {
			if value != rValue {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func AssertKeys(item map[string]interface{}, expect utils.Expect) bool {
	for _, key := range expect.Keys {
		if _, ok := item[key]; !ok {
			return false
		}
	}
	return true
}
