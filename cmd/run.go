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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/prawf/prawf-cli/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a test defined in the prawf.json file",
	Long: `Run a test defined in your prawf.json file.

Runs the test marked as 'current' by default.`,
	Args: cobra.NoArgs,
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
			MakeRequest(test.URL, method.Path, method.Method, method.Header, method.Query, method.Body, method.Name)
		}
	},
}

// MakeRequest makes an HTTP request
func MakeRequest(
	url string,
	path string,
	method string,
	header map[string]interface{},
	query map[string]interface{},
	body map[string]interface{},
	name string) {
	// Create a new HTTP request
	req := NewRequest(url, path, method, header, query, body, name)

	// Create a client and make the request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.WithFields(log.Fields{
			"status code": resp.Status,
		}).Info()
	} else {
		log.WithFields(log.Fields{
			"status code": resp.Status,
		}).Error()
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
	data := string(b)

	log.Info(data)
}

// newRequestBody creates a new request body from the given interface
func newRequestBody(body map[string]interface{}) (*bytes.Buffer, error) {
	bodyJSON, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	bodyRequest := bytes.NewBuffer(bodyJSON)

	return bodyRequest, nil
}

// unmarshalJSON coverts the JSON response to a map
// func unmarshalJSON(b []byte) (map[string]interface{}, error) {
// 	var data map[string]interface{}

// 	err := json.Unmarshal(b, &data)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

func NewRequest(
	url string,
	path string,
	method string,
	header map[string]interface{},
	query map[string]interface{},
	body map[string]interface{},
	name string) *http.Request {
	// Create a request body
	bodyRequest, err := newRequestBody(body)

	if err != nil {
		log.Fatal(err)
	}
	// Create a request with the specified method to the specified url with the created body
	req, err := http.NewRequest(strings.ToUpper(method), url+path, bodyRequest)

	if err != nil {
		log.Fatal(err)
	}
	// Add all the specified headers
	for key, value := range header {
		v := fmt.Sprintf("%v", value)
		req.Header.Add(key, v)
	}
	// Add all the specified queries
	q := req.URL.Query()

	for key, value := range query {
		v := fmt.Sprintf("%v", value)
		q.Add(key, v)
	}

	req.URL.RawQuery = q.Encode()

	log.WithFields(log.Fields{
		"name":   name,
		"method": strings.ToUpper(method),
		"path":   path,
	}).Info()

	log.WithFields(log.Fields{
		"header": header,
		"query":  query,
		"body":   body,
	}).Debug()

	return req
}
