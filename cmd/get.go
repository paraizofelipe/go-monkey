package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/paraizofelipe/go-monkey/api"

	"github.com/spf13/cobra"
)

type Table struct {
	Header []string
	Body   [][]string
}

func IdToShortId(id string) string {
	idShards := strings.Split(id, "-")
	return idShards[len(idShards)-1]
}

func ShowTable(title []string, data []api.Entity) error {
	var header string
	var body string

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, h := range title {
		header += fmt.Sprintf("%s\t", h)
	}
	if _, err := fmt.Fprintln(w, header); err != nil {
		return err
	}

	for _, elem := range data {
		for _, h := range title {
			body += fmt.Sprintf("%v\t", elem.GetValue(h))
		}
		if _, err := fmt.Fprintln(w, body); err != nil {
			return err
		}
		body = ""
	}

	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}

	return nil
}

var getCmd = &cobra.Command{
	Use:   "get [OPTIONS]",
	Short: "get",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Invalid args")
		}

		if _, ok := kong.EndPoints[args[0]]; !ok {
			return errors.New("Invalid args")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "services" {
			err, services := kong.ListServices()
			if err != nil {
				log.Fatal(err)
			}

			ss := make([]api.Entity, len(services))
			for index, _ := range services {
				ss[index] = &services[index]
			}

			header := []string{"ID", "NAME", "PROTOCOL", "PORT"}
			err = ShowTable(header, ss)
			if err != nil {
				log.Fatal(err)
			}
		}

		if args[0] == "routes" {
			err, routes := kong.ListRoutes()
			if err != nil {
				log.Fatal(err)
			}

			//err, r := routes[0].ToMap()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//fmt.Println(r)

			rts := make([]api.Entity, len(routes))
			for index, _ := range routes {
				rts[index] = &routes[index]
			}

			header := []string{"ID", "SERVICE", "NAME", "PROTOCOLS", "METHODS", "PATHS"}
			err = ShowTable(header, rts)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
