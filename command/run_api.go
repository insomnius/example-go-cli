package command

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// NewRunAPI return command run API
func NewRunAPI(engine *gin.Engine) *cobra.Command {
	return &cobra.Command{
		Use:   "run:api",
		Short: "Run API",
		Run: func(cmd *cobra.Command, args []string) {
			engine.Run(":8080")
		},
	}
}
