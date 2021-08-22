package main

import (
	"io"
	"os"
	"time"

	"__anonymous__/__goapp__/src/app"
	"__anonymous__/__goapp__/src/commands"
	"__anonymous__/__goapp__/src/config"
	"__anonymous__/__goapp__/src/http"

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
	// {{end}}
	// {{if eq .database "mongo"}}
	app.SetupMongoDB()
	ctx, cancel := app.MongoOperationCtx()
	defer app.MongoClient().Disconnect(ctx)
	defer cancel()
	// {{end}}
	// {{if eq .web "y"}}
	app.SetupWeb()
	app.Server().Use(recover.New())
	app.Server().Use(middlewares.AccessLogger(createAccessLogger()))
	app.Server().Use(middlewares.Maintenance(app.Cache()))
	app.Server().Use(middlewares.RateLimiter("GLOBAL-LIMITER", 60, 1*time.Minute, app.Cache()))
	http.RegisterRoutes(app.Server())
	app.Server().Static("/", "./public")
	app.CLI().AddCommand(commands.ServeCommand)
	// {{end}}

	// Run App
	app.Run()
}

func createAccessLogger() logger.Logger {
	writers := make([]io.Writer, 1)
	writers[0] = logger.NewFileLogger("./.logs/access", "// {{.name}}", "2006-01-02", app.DateFormatter())
	if !app.Config().Bool("prod", false) {
		writers = append(writers, os.Stdout)
	}

	return logger.NewLogger("2006-01-02 15:04:05", app.DateFormatter(), writers...)
}
