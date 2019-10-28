package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var routeCmd = &cobra.Command{
	Use:   "routes [command]",
	Short: "routes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service")
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)
	serviceCmd.AddCommand(createCmd)
}
