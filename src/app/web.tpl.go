package app

import (
	"github.com/bopher/cache"
	"github.com/bopher/cliutils/maintenance"
	httput "github.com/bopher/http"
	"github.com/bopher/logger"
	"github.com/gofiber/fiber/v2"
)

// SetupWeb driver
func SetupWeb() {
	conf := fiber.Config{}
	conf.DisableStartupMessage = Config().Bool("prod", false)
	conf.ErrorHandler = httput.ErrorLogger(logger.NewLogger(
		"2006-01-02 15:04:05",
		DateFormatter(),
		logger.NewFileLogger(
			"./.logs/error",
			"error",
			"2006-01-02",
			DateFormatter(),
		),
	), DateFormatter(), Config().Bool("prod", false))
	server := fiber.New(conf)
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
