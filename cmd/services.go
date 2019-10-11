package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	endpoint = "/services"
)

type Service struct {
	Name           string `json:"name"`
	Protocol       string `json:"protocol,omitempty"`
	Host           string `json:"host"`
	Port           int64  `json:"port"`
	Path           string `json:"path,omitempty"`
	Retries        string `json:"retries,omitempty"`
	ConnectTimeout string `json:"connect_timeout,omitempty"`
	WriteTimeout   string `json:"write_timeout,omitempty"`
	ReadTimeout    string `json:"read_timeout,omitempty"`
	Url            string `json:"url,omitempty"`
}

func (ks *Service) InfoToTable() string {
	return ""
}

func (ks *Service) ListInfo() string {
	configKong := viper.Get("kong.host").([]interface{})
	kongUrl := configKong[0].(map[string]interface{})["url"].(string)
	u, err := url.Parse(kongUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, endpoint)

	resp, err := http.Get(u.String())
	if err != nil || resp.StatusCode != 200 {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
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
