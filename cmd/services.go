package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "services [command]",
	Short: "services",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service")
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(createCmd)
}
