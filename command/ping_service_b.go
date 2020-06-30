package command

import (
	"fmt"
	"io/ioutil"

	"github.com/dghubble/sling"
	"github.com/spf13/cobra"
)

// NewPingServiceB return command to ping service b availability
func NewPingServiceB(serviceB *sling.Sling) *cobra.Command {
	return &cobra.Command{
		Use:   "ping:service_b",
		Short: "Ping service b.",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := serviceB.Get("/status").Receive(nil, nil)
			if err != nil {
				panic(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(body))
		},
	}
}
