package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewScriptAddToCart return command to do add to cart based on user id and product id
func NewScriptAddToCart() *cobra.Command {
	return &cobra.Command{
		Use:     "script:add_to_cart",
		Short:   "Add to cart based on user_id and product_id",
		Example: "our_service script:add_to_cart [user_id] [product_id]",
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			// You may called logic to your service here
			fmt.Printf("Add to cart for product with id: %s to user with id of -> %s\n", args[0], args[1])
		},
	}
}
