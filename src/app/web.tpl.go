package app

import (
	"github.com/bopher/cache"
	"github.com/bopher/cliutils/maintenance"
	httput "github.com/bopher/http"
	"github.com/gofiber/fiber/v2"
)

// SetupWeb driver
func SetupWeb() {
	conf := confOrPanic()
	onProd := conf.Cast("prod").BoolSafe(true)
	erLogger := Logger("error_logger")
	if erLogger == nil {
		panic("failed to find error_logger")
	}

	server := fiber.New(fiber.Config{
		DisableStartupMessage: onProd,
		ErrorHandler:          httput.ErrorLogger(erLogger, DateFormatter(), onProd),
	})
	_container.Register("--APP-SERVER", server)

	_cli.AddCommand(maintenance.DownCommand(func(driver string) cache.Cache {
		return Cache(driver)
	}, "--APP-CACHE"))
	_cli.AddCommand(maintenance.UpCommand(func(driver string) cache.Cache {
		return Cache(driver)
	}, "--APP-CACHE"))
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
