package app

import (
	"runtime/debug"
	"time"

	"github.com/bopher/cli"
	"github.com/bopher/container"
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

// Resolve get app dependency
func Resolve[T any](name string, fallback T) (T, bool) {
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(T); ok {
			return res, true
		}
	}
	return fallback, false
}

// CLI get app main cli
func CLI() cli.CLI {
	return _cli
}

// Run cli and log panic if exists
func Run() {
	defer func() {
		if r := recover(); r != nil {
			logger := Logger("crash_logger")
			if logger != nil {
				logger.Divider("=", 100, "APP CRASHED")
				logger.Error().Print("%v", r)
				logger.Raw("\n\nStacktrace:\n")
				logger.Raw(string(debug.Stack()))
				logger.Divider("=", 100, DateFormatter()(time.Now().UTC(), "2006-01-02 15:04:05"))
				logger.Raw("\n\n")
			}
		}
	}()
	_cli.Run()
}
