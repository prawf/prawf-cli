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
	"fmt"
	"log"

	"github.com/prawf/prawf-cli/utils"

	"github.com/spf13/cobra"
)

var (
	configPath string
	configName string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create prawf.json file",
	Long: `Create and initialise prawf.json file in the current working directory.

The file will be initialised with default values.`,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig(configPath, configName)

	},
}

func init() {

	cp, err := utils.GetDir()
	if err != nil {
		log.Fatal(err)
	}

	initCmd.Flags().StringVarP(&configPath, "path", "p", cp, "path to create the prawf.json file")
	initCmd.Flags().StringVarP(&configName, "name", "n", utils.ConfigName, "name of the config file")
}

func initConfig(cp string, cn string) {
	conf := utils.GetConf(cp, cn)

	fmt.Printf("Hostname: %v\n", conf.Hostname)
	fmt.Printf("Port: %v\n", conf.Port)
}
