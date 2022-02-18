package main

import (
	"os"

	"mekramy/__boiler/src/app"
	"mekramy/__boiler/src/commands"
	"mekramy/__boiler/src/config"

	// {{if eq .web "y"}}
	"mekramy/__boiler/src/http"

	"github.com/bopher/http/middlewares"
	"github.com/bopher/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// {{end}}
	// {{if eq .database "mysql"}}
	"github.com/bopher/database/migration"
	// {{end}}
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
	app.SetupWeb(http.OnError)
	app.Server().Use(recover.New())
	if app.Config().Cast("web.log").BoolSafe(false) {
		onDev := !app.Config().Cast("prod").BoolSafe(true)
		appName := app.Config().Cast("name").StringSafe("// {{.name}}")
		_logger := logger.NewLogger("2006-01-02 15:04:05", app.DateFormatter())
		_logger.AddWriter("main", logger.NewFileLogger(app.LogPath("access"), appName, "2006-01-02", app.DateFormatter()))
		if onDev {
			_logger.AddWriter("dev", os.Stdout)
		}
		app.Server().Use(middlewares.AccessLogger(_logger))
	}
	http.RegisterGlobalMiddlewares(app.Server())
	http.RegisterRoutes(app.Server())
	app.Server().Static("/", "./public")
	// {{end}}

	// Register commands and run app
	app.CLI().AddCommand(commands.HashCommand(app.CryptoResolver, "--APP-CRYPTO"))
	app.CLI().AddCommand(commands.ClearCommand)
	app.CLI().AddCommand(commands.DownCommand)
	app.CLI().AddCommand(commands.UpCommand)
	app.CLI().AddCommand(commands.VersionCommand)
	// {{if eq .database "mysql"}}
	app.CLI().AddCommand(migration.MigrationCommand(app.DatabaseResolver, "--APP-DB", "./database/migrations", "./database/seeds"))
	// {{end}}
	// {{if eq .web "y"}}
	app.CLI().AddCommand(commands.ServeCommand)
	// {{end}}
	app.Run()
}
