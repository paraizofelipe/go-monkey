package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/paraizofelipe/go-monkey/api"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var a *api.Api

func IdToShortId(id string) string {
	idShards := strings.Split(id, "-")
	return idShards[len(idShards)-1]
}

func ShowTable() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	header := fmt.Sprintf("%s\t%s\t%s", "ID", "HOST", "NAME")

	_, err := fmt.Fprintln(w, header)
	if err != nil {
		log.Fatal(err)
	}

	for _, svc := range s.ListInfo() {
		body := fmt.Sprintf("%s\t%s\t%s", IdToShortId(svc.ID), svc.Host, svc.Name)
		if _, err := fmt.Fprintln(w, body); err != nil {
			log.Fatal(err)
		}
	}

	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}

var getCmd = &cobra.Command{
	Use:   "get [OPTIONS]",
	Short: "get",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Invalid args")
		}

		if _, ok := a.EndPoints[args[0]]; !ok {
			return errors.New("Invalid args")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		configKong := viper.Get("kong.host").([]interface{})
		baseUrl := configKong[0].(map[string]interface{})["url"].(string)

		a = api.New(baseUrl)
		if args[0] == "service" {
			err, services := a.ListServices()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(services)
		}

		if args[0] == "routes" {
			err, routes := a.ListRoutes()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(routes)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
