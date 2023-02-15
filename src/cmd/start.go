package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.sima-land.ru/dev-dep/secretary-bot/app"
)

func start(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start TGBot",
		Run: func(cmd *cobra.Command, args []string) {
			app.Serve()
		},
	}
}
