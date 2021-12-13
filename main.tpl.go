package main

import (
	"os"
	"time"

	"mekramy/__boiler/src/app"
	"mekramy/__boiler/src/commands"
	"mekramy/__boiler/src/config"
	"mekramy/__boiler/src/http"

	"github.com/bopher/http/middlewares"
	"github.com/bopher/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app.SetupConfig()
	config.Configure(app.Config())
	app.SetupCache()
	app.SetupCrypto()
	app.SetupLogger()
	app.SetupTranslator()
	config.ConfigureMessages(app.Translator())
	app.SetupValidator()
	// {{if eq .database "mysql"}}
	app.SetupDatabase()
	defer app.Database().Close()
	// {{end}}
	// {{if eq .database "mongo"}}
	app.SetupMongoDB()
	ctx, cancel := app.MongoOperationCtx()
	defer cancel()
	defer app.MongoClient().Disconnect(ctx)
	// {{end}}
	// {{if eq .web "y"}}
	limit := app.Config().Cast("web.limit").UInt32Safe(60)
	log_access := app.Config().Cast("web.log").BoolSafe(false)

	app.SetupWeb()
	app.Server().Use(recover.New())
	if log_access {
		onDev := !app.Config().Cast("prod").BoolSafe(true)
		appName := app.Config().Cast("name").StringSafe("// {{.name}}")
		_logger := logger.NewLogger("2006-01-02 15:04:05", app.DateFormatter())
		_logger.AddWriter("main", logger.NewFileLogger(app.LogPath("access"), appName, "2006-01-02", app.DateFormatter()))
		if onDev {
			_logger.AddWriter("dev", os.Stdout)
		}
		app.Server().Use(middlewares.AccessLogger(_logger))
	}
	app.Server().Use(middlewares.Maintenance(app.Cache()))
	app.Server().Use(middlewares.RateLimiter("GLOBAL-LIMITER", limit, 1*time.Minute, app.Cache()))
	http.RegisterRoutes(app.Server())
	app.Server().Static("/", "./public")
	app.CLI().AddCommand(commands.ServeCommand)
	// {{end}}

	// Run App
	app.CLI().AddCommand(commands.VersionCommand)
	app.Run()
}
