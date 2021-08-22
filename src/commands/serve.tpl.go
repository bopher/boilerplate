package commands

import (
	"__anonymous__/__goapp__/src/app"

	"github.com/spf13/cobra"
)

// ServeCommand serve web app
var ServeCommand = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run: func(cmd *cobra.Command, args []string) {
		app.Server().Listen(app.Config().String("web.port", ":8888"))
	},
}
