package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func containsCommand(name string, commands []*cobra.Command) bool {
	for _, command := range commands {
		if command.Name() == name {
			return true
		}
	}
	return false
}

var publishedCmds map[string]interface{}

var getCmd = &cobra.Command{
	Use:   "get [OPTIONS]",
	Short: "get",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Invalid args")
		}

		if _, ok := publishedCmds[args[0]]; !ok {
			return errors.New("Invalid args")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		service := publishedCmds[args[0]].(Service)
		fmt.Println(service.ListInfo())
	},
}

func init() {
	publishedCmds = map[string]interface{}{"services": Service{}}
	rootCmd.AddCommand(getCmd)
}
