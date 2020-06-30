package command

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

// NewPingDB return command run to check mysql availability
func NewPingDB(db *sql.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "ping:mysql",
		Short: "Check mysql availability",
		Run: func(cmd *cobra.Command, args []string) {
			if err := db.Ping(); err != nil {
				fmt.Println("Failed to ping mysql databases:", err)
				return
			}

			fmt.Println("Successfully ping mysql databases")
		},
	}
}
