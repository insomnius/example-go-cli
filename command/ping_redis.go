package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/redis.v5"
)

// NewPingRedis return command run to check redis availability
func NewPingRedis(rdb *redis.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "ping:redis",
		Short: "Check redis availability",
		Run: func(cmd *cobra.Command, args []string) {
			pong, err := rdb.Ping().Result()

			fmt.Println(pong, err)
		},
	}
}
