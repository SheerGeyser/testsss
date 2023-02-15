package cmd

import (
	"fmt"

	"gitlab.sima-land.ru/dev-dep/secretary-bot/app"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := cobra.Command{
		Use:   "secretary-bot",
		Short: "TGBot application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Type a 'secretary-bot help' for usage details")
		},
	}

	root.AddCommand(
		start(app.NewApp(app.Config{})),
	)

	return &root
}
