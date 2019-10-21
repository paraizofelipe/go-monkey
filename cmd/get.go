package cmd

import (
	"errors"
	"log"
	"net/url"
	"path"
	"strings"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

type CmdReader interface {
	ShowTable()
}

var publishedCmds map[string]CmdReader

func IdToShortId(id string) string {
	idShards := strings.Split(id, "-")
	return idShards[len(idShards)-1]
}

func getInfo(endpoint string) {
	configKong := viper.Get("kong.host").([]interface{})
	kongUrl := configKong[0].(map[string]interface{})["url"].(string)
	u, err := url.Parse(kongUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, endpoint)
}

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
		service := publishedCmds[args[0]].(CmdReader)
		service.ShowTable()
	},
}

func init() {
	publishedCmds = make(map[string]CmdReader)
	publishedCmds["services"] = new(Service)
	publishedCmds["routes"] = new(Route)
	rootCmd.AddCommand(getCmd)
}
