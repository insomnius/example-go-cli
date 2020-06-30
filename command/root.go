package command

import "github.com/spf13/cobra"

// NewRoot return root command of our application
func NewRoot() *cobra.Command {
	return &cobra.Command{
		Use:   "our_service",
		Short: "Example CLI command in Go using cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
}
