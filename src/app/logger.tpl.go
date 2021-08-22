package app

import (
	"io"
	"os"

	cstorage "github.com/bopher/cliutils/storage"
	"github.com/bopher/logger"
)

// SetupLogger driver
func SetupLogger() {
	writers := make([]io.Writer, 0)
	writers = append(writers, logger.NewFileLogger("./.logs/log", "// {{.name}}", "2006-01-02", DateFormatter()))
	if !Config().Bool("prod", false) {
		writers = append(writers, os.Stdout)
	}
	if l := logger.NewLogger("2006-01-02 15:04:05", DateFormatter(), writers...); l != nil {
		_container.Register("--APP-LOGGER", l)
	} else {
		panic("failed to build crypto driver")
	}

	_cli.AddCommand(cstorage.ClearCommand("./.logs"))
}

// Logger get logger driver
// leave name empty to resolve default
func Logger(names ...string) logger.Logger {
	name := "--APP-LOGGER"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(logger.Logger); ok {
			return res
		}
	}
	return nil
}

// DateFormatter get default app date formatter
func DateFormatter() logger.TimeFormatter {
	// {{if eq .locale "fa"}}
	return logger.JalaliFormatter
	// {{else}}
	return logger.GregorianFormatter
	// {{end}}
}
