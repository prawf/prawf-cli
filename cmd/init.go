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
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create prawf.json file",
	Long: `Create and initialise prawf.json file in the current working directory.

The file will be initialised with default values.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// This is an empty command, running this when the user first initialises will
		// call loadConfig() function in root.go which will initialise the current directory
		// with a prawf.json file
	},
}
