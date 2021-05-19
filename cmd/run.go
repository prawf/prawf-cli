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
	"log"
	"net/http"

	"github.com/prawf/prawf-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a test defined in the prawf.json file",
	Long: `Run a test defined in your prawf.json file.

Runs the test marked as 'current-test' by default.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := utils.GetPrawfConfig(viper.GetViper())

		if err != nil {
			log.Fatal(err)
		}

		test, err := conf.GetTest()

		if err != nil {
			log.Fatal(err)
		}

		for _, method := range test.Methods {
			MakeRequest(test.URL, method.Path, method.Method, method.Header, method.Query, method.Body)
		}
	},
}

func MakeRequest(
	url string,
	path string,
	method string,
	header map[string]interface{},
	query map[string]interface{},
	body map[string]interface{}) {

	bodyJSON, err := json.Marshal(body)
	bodyRequest := bytes.NewBuffer(bodyJSON)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(method, url+path, bodyRequest)

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range header {
		v := fmt.Sprintf("%v", value)
		req.Header.Add(key, v)
	}

	for key, value := range query {
		v := fmt.Sprintf("%v", value)
		req.Header.Add(key, v)
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data interface{}

	err = json.Unmarshal(b, &data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)

}
