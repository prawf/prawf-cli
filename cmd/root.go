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

	"github.com/prawf/prawf-cli/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Users can specify their own path and name to create the prawf.json file
var (
	configPath string
	configName string
	outputFmt  string
)

var rootCmd = &cobra.Command{
	Use:   "prawf",
	Short: "API Testing made easy",
	Long: `prawf is an all in one API testing framework.
	
It lets you define and run tests on your API endpoints and/or make individual requests to your endpoints.`,
}

func init() {
	// Check if the passed in output flag is valid
	cobra.OnInitialize(validateOutputFlag)
	// Load the users config file (prawf.json) file when a command is run
	cobra.OnInitialize(loadConfig)
	// Set the default directory as the current directory
	cp, err := utils.GetDir()
	if err != nil {
		log.Fatal(err)
	}
	// All commands will have flags to pass in the path to and the name of the config file
	// By default, these are the current directory and prawf respectievely
	rootCmd.PersistentFlags().StringVarP(&configPath, "path", "p", cp, "path to create the prawf.json file")
	rootCmd.PersistentFlags().StringVarP(&configName, "name", "n", utils.ConfigName, "name of the config file")
	rootCmd.PersistentFlags().StringVarP(&outputFmt, "output", "o", "", "format output (available values: [json])")
	// Initialise the commands
	commands := []*cobra.Command{
		initCmd,
		runCmd,
		reqCmd,
	}

	rootCmd.AddCommand(commands...)
	// Remove the unnecessary `help` command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Execute() {
	rootCmd.Execute()
}

// loadConfig loads the users config file (prawf.json) if it exists or initialises a new one if it doesn't
func loadConfig() {
	filePath, fileName := utils.GetFilePath(configPath, configName)

	if !utils.FileExists(filePath) {
		if utils.AskForConfirmation(fmt.Sprintf("%s file not found. Would you like to create one?", fileName)) {
			initConfig(filePath, fileName)
		} else {
			log.WithField("file", fileName).Fatalf(
				"File not found. Run `prawf init -p %s -n %s` to create one.",
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

	log.WithField("file", fileName).Info("File loaded.")
}

// initConfig creates a new config file with in the specified path with the specified name
func initConfig(filePath string, fileName string) {
	err := utils.CreateConfigFile(filePath)

	if err != nil {
		log.Fatal(err)
	}
	// Initialise the config file with a sample test definition
	err = utils.AddTestsToConfig("sample-test", utils.TemplateTest, filePath)

	if err != nil {
		log.Fatal(err)
	}

	log.WithField("file", fileName).Info("File created.")
}

func validateOutputFlag() {
	if outputFmt != "" {
		if outputFmt == "json" {
			log.SetFormatter(&log.JSONFormatter{})
			return
		}
		log.WithField("output", outputFmt).Fatal("Invalid output format specified.")
	}
}
