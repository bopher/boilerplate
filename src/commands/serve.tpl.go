package commands

import (
	"fmt"
	"mekramy/__boiler/src/app"

	"github.com/spf13/cobra"
)

// ServeCommand serve web app
var ServeCommand = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run: func(cmd *cobra.Command, args []string) {
		app.Server().Listen(fmt.Sprintf(":%d", app.Config().Cast("web.port").IntSafe(8888)))
	},
}
