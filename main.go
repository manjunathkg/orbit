/*
Package main is the root package of the application.

Orbit is a cross-platform task runner for executing commands and generating files from templates.

It started with the need to find a cross-platform alternative of "make"
and "sed -i" commands. As it does not aim to be as powerful as these two
commands, Orbit offers an elegant solution for running tasks and generating
files from templates, whatever the platform you're using.

For more information, go to https://github.com/manjunathkg/orbit.
*/
package main

import (
	"os"

	"github.com/manjunathkg/orbit/app"
	"github.com/manjunathkg/orbit/app/logger"
	OrbitVersion "github.com/manjunathkg/orbit/app/version"
)

/*
version will be set by GoReleaser.

It will be the current Git tag (with v prefix stripped) or
the name of the snapshot if you're using the --snapshot flag.
*/
var version = "master"

// main is the root function of the application.
func main() {
	OrbitVersion.Current = version

	if err := app.RootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
