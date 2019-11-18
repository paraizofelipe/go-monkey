package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

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

func ShowTable(table Table) error {
	var header string
	var body string

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, h := range table.Header {
		header += fmt.Sprintf("%s\t", h)
	}
	if _, err := fmt.Fprintln(w, header); err != nil {
		return err
	}

	for _, elem := range table.Body {
		row := strings.Join(elem, "\t")
		body += fmt.Sprintf("%s", row)
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

			header := []string{"ID", "NAME", "PROTOCOL", "PORT"}
			body := make([][]string, len(services))
			for row, r := range services {
				body[row] = append(body[row], IdToShortId(r.Id))
				body[row] = append(body[row], r.Name)
				body[row] = append(body[row], r.Protocol)
				body[row] = append(body[row], fmt.Sprintf("%d", r.Port))
			}

			table := Table{
				header,
				body,
			}
			err = ShowTable(table)
			if err != nil {
				log.Fatal(err)
			}
		}

		if args[0] == "consumers" {
			err, consumers := kong.ListConsumer()
			if err != nil {
				log.Fatal(err)
			}

			header := []string{"ID", "USERNAME", "CUSTOM_ID", "TAGS"}
			body := make([][]string, len(consumers))
			for row, c := range consumers {
				body[row] = append(body[row], IdToShortId(c.Id))
				body[row] = append(body[row], c.Username)
				body[row] = append(body[row], c.CustomId)
				body[row] = append(body[row], fmt.Sprintf("%v", c.Tags))
			}

			table := Table{
				header,
				body,
			}
			err = ShowTable(table)
			if err != nil {
				log.Fatal(err)
			}
		}

		if args[0] == "routes" {
			err, routes := kong.ListRoutes()
			if err != nil {
				log.Fatal(err)
			}

			header := []string{"ID", "SERVICE", "NAME", "PROTOCOLS", "METHODS", "PATHS"}
			body := make([][]string, len(routes))
			for row, r := range routes {
				body[row] = append(body[row], IdToShortId(r.Id))
				body[row] = append(body[row], r.Service.Id)
				body[row] = append(body[row], r.Name)
				body[row] = append(body[row], fmt.Sprintf("%s", r.Protocols))
				body[row] = append(body[row], fmt.Sprintf("%s", r.Methods))
				body[row] = append(body[row], fmt.Sprintf("%s", r.Paths))
			}

			table := Table{
				header,
				body,
			}
			err = ShowTable(table)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
