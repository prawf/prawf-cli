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
	log.Info(string(b))

	// Get the json data
	err = json.Unmarshal([]byte(b), &data)

	if err != nil {
		log.Fatal(err)
	}

	// The response could be an interface or an array of interface
	if m, ok := data.([]interface{}); ok {
		for _, item := range m {
			Expect(item.(map[string]interface{}))
		}
	} else if item, ok := data.(map[string]interface{}); ok {
		Expect(item)
	} else {
		log.Fatal("Invalid response received.")
	}
}

func Expect(item map[string]interface{}) {
	// for key, value := range item {
	// 	log.Info(key, value)
	// }
}
