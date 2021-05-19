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
	"github.com/spf13/viper"
)

var (
	configPath string
	configName string
)

var rootCmd = &cobra.Command{
	Use:   "prawf",
	Short: "API Testing made easy",
	Long: `prawf is an all in one API testing framework.
	
It lets you define and run tests on your API endpoints and/or make individual requests to your endpoints.`,
}

func init() {

	cobra.OnInitialize(loadConfig)

	cp, err := utils.GetDir()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(&configPath, "path", "p", cp, "path to create the prawf.json file")
	rootCmd.PersistentFlags().StringVarP(&configName, "name", "n", utils.ConfigName, "name of the config file")

	commands := []*cobra.Command{
		initCmd,
		runCmd,
		reqCmd,
	}

	rootCmd.AddCommand(commands...)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Execute() {
	rootCmd.Execute()
}

func loadConfig() {
	filePath, fileName := utils.GetFilePath(configPath, configName)

	if !utils.FileExists(filePath) {
		if utils.AskForConfirmation(fmt.Sprintf("%s file not found. Would you like to create one?", fileName)) {
			initConfig(filePath, fileName)
		} else {
			log.Fatalf(
				"%s not found. Run `prawf init -p %s -n %s` to create one.",
				fileName,
				configPath,
				configName,
			)
		}
	}

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s loaded.", fileName)
}

func initConfig(filePath string, fileName string) {
	err := utils.CreateConfigFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	err = utils.AddTestsToConfig("sample-test", utils.TemplateTest, filePath)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s created.", fileName)
}
