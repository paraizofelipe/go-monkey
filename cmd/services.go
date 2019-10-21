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

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	serviceEndpoint = "/services"
)

type Service struct {
	Host           string      `json:"host"`
	CreatedAt      int64       `json:"created_at"`
	ConnectTimeout int64       `json:"connect_timeout"`
	ID             string      `json:"id"`
	Protocol       string      `json:"protocol"`
	Name           string      `json:"name"`
	ReadTimeout    int64       `json:"read_timeout"`
	Port           int64       `json:"port"`
	Path           interface{} `json:"path"`
	UpdatedAt      int64       `json:"updated_at"`
	Retries        int64       `json:"retries"`
	WriteTimeout   int64       `json:"write_timeout"`
}

type RespService struct {
	Next interface{} `json:"next"`
	Data []Service   `json:"data"`
}

func (s *Service) ShowTable() {
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

func (s *Service) ListInfo() []Service {
	configKong := viper.Get("kong.host").([]interface{})
	kongUrl := configKong[0].(map[string]interface{})["url"].(string)
	u, err := url.Parse(kongUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, serviceEndpoint)

	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var services RespService
	err = json.Unmarshal(body, &services)
	if err != nil {
		log.Fatal(err)
	}

	return services.Data
}

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
