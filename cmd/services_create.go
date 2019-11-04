package cmd

import (
	"errors"
	"log"
	"net/url"
	"strconv"

	"github.com/paraizofelipe/go-monkey/api"
	"github.com/spf13/cobra"
)

var sc api.Service

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS]",
	Short: "Create service",
	Long:  `Create service in API Gateway Kong`,
	Args: func(cmd *cobra.Command, args []string) error {
		var err error

		if len(args) < 2 {
			return errors.New("Invalid args")
		}

		u, err := url.Parse(args[1])
		if err != nil {
			return errors.New("Invalid arg to host")
		}

		sc.Name = args[0]
		sc.Host = u.Hostname()
		sc.Port, err = strconv.ParseInt(u.Port(), 10, 64)
		if err != nil {
			return errors.New("")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := kong.CreateServices(sc)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVar(&sc.Protocol, "protocol", "", `The protocol used to communicate with the upstream. It can be one of http (default) or https.`)
	//createCmd.PersistentFlags().StringVar(&sc.Url, "url", "", "Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never “returns” the url).")
	//createCmd.PersistentFlags().StringVar(&sc.Path, "path", "", "The path to be used in requests to the upstream server. Empty by default.")
	createCmd.PersistentFlags().Int64Var(&sc.Retries, "retries", 0, "The number of retries to execute upon failure to proxy. The default is 5.")
	createCmd.PersistentFlags().Int64Var(&sc.ConnectTimeout, "connect-timeout", 0, "The timeout in milliseconds for establishing a connection to the upstream server. Defaults to 60000.")
	createCmd.PersistentFlags().Int64Var(&sc.WriteTimeout, "write-timeout", 0, "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server. Defaults to 60000.")
	createCmd.PersistentFlags().Int64Var(&sc.ReadTimeout, "read-timeout", 0, "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server. Defaults to 60000.")

	serviceCmd.AddCommand(createCmd)
}
