package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const (
	routesEndpoint = "/routes"
)

type Route struct {
	ID                      string              `json:"id"`
	CreatedAt               int64               `json:"create_at"`
	UpdateAt                int64               `json:"update_at"`
	Name                    string              `json:"name"`
	Protocols               []string            `json:"protocols"`
	Methods                 []string            `json:"methods"`
	Hosts                   []string            `json:"hosts"`
	Paths                   []string            `json:"paths"`
	Headers                 map[string][]string `json:"headers"`
	HttpsRedirectStatusCode int64               `json:"https_redirect_status_code"`
	RegexPriority           int64               `json:"regex_priority"`
	StripPath               bool                `json:"strip_path"`
	PreserveHost            bool                `json:"preserve_host"`
	Tags                    string              `json:"tags"`
	Service                 map[string]string   `json:"service"`
}

type RespRoute struct {
	Next interface{} `json:"next"`
	Data []Route     `json:"data"`
}

func (r *Route) ShowTable() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	header := fmt.Sprintf("%s\t%s\t%s\t%s\t%s", "ID", "HOSTS", "PATHS", "SERVICE", "NAME")

	_, err := fmt.Fprintln(w, header)
	if err != nil {
		log.Fatal(err)
	}

	for _, rts := range r.ListInfo() {
		body := fmt.Sprintf("%s\t%s\t%s\t%s\t%s", IdToShortId(rts.ID), rts.Hosts, rts.Paths, IdToShortId(rts.Service["id"]), rts.Name)
		if _, err := fmt.Fprintln(w, body); err != nil {
			log.Fatal(err)
		}
	}

	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}

func (r *Route) ListInfo() []Route {
	configKong := viper.Get("kong.host").([]interface{})
	kongUrl := configKong[0].(map[string]interface{})["url"].(string)
	u, err := url.Parse(kongUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, routesEndpoint)

	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var routes RespRoute
	err = json.Unmarshal(body, &routes)
	if err != nil {
		log.Fatal(err)
	}

	return routes.Data
}

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
