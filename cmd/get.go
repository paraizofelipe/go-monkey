package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/paraizofelipe/go-monkey/api"

	"github.com/spf13/cobra"
)

func IdToShortId(id string) string {
	idShards := strings.Split(id, "-")
	return idShards[len(idShards)-1]
}

func StructToRow(s interface{}) (error, []string) {
	var rows []string
	var inInterface map[string]interface{}

	inrec, _ := json.Marshal(s)
	if err := json.Unmarshal(inrec, &inInterface); err != nil {
		return err, nil
	}

	for v := range inInterface {
		rows = append(rows, v)
	}

	return nil, rows
}

func Show(data []api.Entity) {
	fmt.Println(data[0].GetValue("Host"))
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

	for index, elem := range data {
		body += fmt.Sprintf("%s\t", elem.GetValue(title[index]))
	}
	if _, err := fmt.Fprintln(w, body); err != nil {
		return err
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
			for i, v := range services {
				ss[i] = &v
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
			fmt.Println(routes)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
