// Package app implements all commands of the application.
package app

import (
	"github.com/manjunathkg/orbit/app/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// templateFilePath is the path of a data-driven template.
	templateFilePath string

	// payload represents a map of YAML files, TOML files, JSON files, .env files and raw data.
	// Value format: key,path;key,path;key,data...
	payload string

	// templates represents a map of additional templates.
	// Value format: path,path,path...
	templates string

	// verbose enables info logs if true.
	verbose bool

	// debug enables debug logs if true.
	debug bool

	// RootCmd is the instance of the root of all commands.
	RootCmd = &cobra.Command{
		Use:           "orbit",
		Short:         "A cross-platform task runner for executing commands and generating files from templates",
		Long:          "A cross-platform task runner for executing commands and generating files from templates.",
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				logger.SetLevel(logrus.InfoLevel)
			}

			if debug {
				logger.SetLevel(logrus.DebugLevel)
			}
		},
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&templateFilePath, "file", "f", "", "specify the path of a data-driven template")
	RootCmd.PersistentFlags().StringVarP(&payload, "payload", "p", "", "specify a map of YAML files, TOML files, JSON files, .env files and raw data")
	RootCmd.PersistentFlags().StringVarP(&templates, "templates", "t", "", "specify a map of additional templates")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set logging to info level")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "set logging to debug level")
}
