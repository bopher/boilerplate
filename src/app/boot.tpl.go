package app

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/bopher/cli"
	"github.com/bopher/container"
	"github.com/bopher/logger"
)

var _container container.Container
var _cli cli.CLI

func init() {
	_container = container.NewContainer()
	_cli = cli.NewCLI("// {{.name}}", "// {{.description}}")
}

// Container get app main container
func Container() container.Container {
	return _container
}

// CLI get app main cli
func CLI() cli.CLI {
	return _cli
}

// Run cli and log panic if exists
func Run() {
	defer func() {
		if r := recover(); r != nil {
			erLogger := logger.NewLogger(
				"2006-01-02 15:04:05",
				DateFormatter(),
				logger.NewFileLogger(
					"./.logs/error",
					"crash",
					"2006-01-02",
					DateFormatter(),
				),
				os.Stdout,
			)
			erLogger.Divider("=", 100, "APP CRASHED")
			erLogger.Error().Print("%v", r)
			erLogger.Raw("\n\nStacktrace:\n")
			erLogger.Raw(string(debug.Stack()))
			erLogger.Divider("=", 100, DateFormatter()(time.Now().UTC(), "2006-01-02 15:04:05"))
			erLogger.Raw("\n\n")
		}
	}()
	_cli.Run()
}
