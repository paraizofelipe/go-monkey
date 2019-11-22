package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/mitchellh/mapstructure"

	"github.com/paraizofelipe/go-monkey/api"

	"github.com/spf13/cobra"
)

type Table struct {
	Header []string
	Body   [][]string
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
		err, entities := kong.ListEntities(args[0])
		if err != nil {
			log.Fatal(err)
		}

		if args[0] == "services" {
			var services []api.Service
			header := []string{"ID", "NAME", "PROTOCOL", "PORT"}
			body := make([][]string, len(entities))

			err := mapstructure.Decode(entities, &services)

			for row, s := range services {
				body[row] = append(body[row], IdToShortId(s.Id))
				body[row] = append(body[row], s.Name)
				body[row] = append(body[row], s.Protocol)
				body[row] = append(body[row], fmt.Sprintf("%d", s.Port))
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

		//if args[0] == "services" {
		//	err, services := kong.Services()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	header := []string{"ID", "NAME", "PROTOCOL", "PORT"}
		//	body := make([][]string, len(services))
		//	for row, r := range services {
		//		body[row] = append(body[row], IdToShortId(r.Id))
		//		body[row] = append(body[row], r.Name)
		//		body[row] = append(body[row], r.Protocol)
		//		body[row] = append(body[row], fmt.Sprintf("%d", r.Port))
		//	}
		//
		//	table := Table{
		//		header,
		//		body,
		//	}
		//	err = ShowTable(table)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}
		//
		//if args[0] == "consumers" {
		//	err, consumers := kong.Consumers()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	header := []string{"ID", "USERNAME", "CUSTOM_ID", "TAGS"}
		//	body := make([][]string, len(consumers))
		//	for row, c := range consumers {
		//		body[row] = append(body[row], IdToShortId(c.Id))
		//		body[row] = append(body[row], c.Username)
		//		body[row] = append(body[row], c.CustomId)
		//		body[row] = append(body[row], fmt.Sprintf("%v", c.Tags))
		//	}
		//
		//	table := Table{
		//		header,
		//		body,
		//	}
		//	err = ShowTable(table)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}
		//
		//if args[0] == "routes" {
		//	err, routes := kong.Routes()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	header := []string{"ID", "SERVICE", "NAME", "PROTOCOLS", "METHODS", "PATHS"}
		//	body := make([][]string, len(routes))
		//	for row, r := range routes {
		//		body[row] = append(body[row], IdToShortId(r.Id))
		//		body[row] = append(body[row], r.Service.Id)
		//		body[row] = append(body[row], r.Name)
		//		body[row] = append(body[row], fmt.Sprintf("%s", r.Protocols))
		//		body[row] = append(body[row], fmt.Sprintf("%s", r.Methods))
		//		body[row] = append(body[row], fmt.Sprintf("%s", r.Paths))
		//	}
		//
		//	table := Table{
		//		header,
		//		body,
		//	}
		//	err = ShowTable(table)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}
		//
		//if args[0] == "upstreams" {
		//	err, upstreams := kong.Upstreams()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	header := []string{"ID", "NAME", "SLOTS", "CREATED"}
		//	body := make([][]string, len(upstreams))
		//	for row, u := range upstreams {
		//		body[row] = append(body[row], IdToShortId(u.Id))
		//		body[row] = append(body[row], u.Name)
		//		body[row] = append(body[row], fmt.Sprintf("%d", u.Slots))
		//		body[row] = append(body[row], fmt.Sprintf("%d", u.CreatedAt))
		//	}
		//
		//	table := Table{
		//		header,
		//		body,
		//	}
		//	err = ShowTable(table)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
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
