package app

import (
	"os"

	"github.com/bopher/logger"
)

// SetupLogger driver
func SetupLogger() {
	conf := confOrPanic()
	appName := conf.Cast("name").StringSafe("// {{.name}}")
	onDev := !conf.Cast("prod").BoolSafe(true)

	// Crash logger
	_crLogger := logger.NewLogger("2006-01-02 15:04:05", DateFormatter())
	_crLogger.AddWriter("main", logger.NewFileLogger(LogPath("crash"), "crash", "2006-01-02", DateFormatter()))
	if onDev {
		_crLogger.AddWriter("dev", os.Stdout)
	}
	_container.Register("crash_logger", _crLogger)

	// Error logger
	_erLogger := logger.NewLogger("2006-01-02 15:04:05", DateFormatter())
	_erLogger.AddWriter("main", logger.NewFileLogger(LogPath("error"), "error", "2006-01-02", DateFormatter()))
	if onDev {
		_erLogger.AddWriter("dev", os.Stdout)
	}
	_container.Register("error_logger", _erLogger)

	// Default logger
	_logger := logger.NewLogger("2006-01-02 15:04:05", DateFormatter())
	_logger.AddWriter("main", logger.NewFileLogger(LogPath("log"), appName, "2006-01-02", DateFormatter()))
	if onDev {
		_logger.AddWriter("dev", os.Stdout)
	}
	_container.Register("--APP-LOGGER", _logger)
}

// Logger get logger driver
// leave name empty to resolve default
// pre defined loggers are crash_logger and error_logger
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
