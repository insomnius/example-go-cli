package command

import (
	"fmt"
	"io/ioutil"

	"github.com/dghubble/sling"
	"github.com/spf13/cobra"
)

// NewPingServiceA return command to ping service a availability
func NewPingServiceA(serviceA *sling.Sling) *cobra.Command {
	return &cobra.Command{
		Use:   "ping:service_a",
		Short: "Ping service a.",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := serviceA.Get("/status").Receive(nil, nil)
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
