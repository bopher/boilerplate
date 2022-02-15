package app

import (
	httput "github.com/bopher/http"
	"github.com/gofiber/fiber/v2"
)

// SetupWeb driver
func SetupWeb(onError httput.ErrorCallback) {
	conf := confOrPanic()
	erLogger := Logger("error_logger")
	if erLogger == nil {
		panic("failed to find error_logger")
	}

	server := fiber.New(fiber.Config{
		DisableStartupMessage: conf.Cast("prod").BoolSafe(true),
		ErrorHandler:          httput.ErrorLogger(erLogger, DateFormatter(), onError),
	})
	_container.Register("--APP-SERVER", server)
}

// Server get web server driver
// leave name empty to resolve default
func Server(names ...string) *fiber.App {
	name := "--APP-SERVER"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := _container.Resolve(name); exists {
		if res, ok := dep.(*fiber.App); ok {
			return res
		}
	}
	return nil
}
