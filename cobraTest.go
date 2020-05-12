package main

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "cobrause",
	Short: "cobra short desc",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("must set URL!")
		}
		client := http.Client{}
		req, err := http.NewRequest("GET". args[], nil)
		if err != nil {
			log.Fatalln("unable to get request")
		}
		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}
	},
}

var username, password string

func main() {
	cmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username for credentials")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for credentials")
	cmd.Execute()
}
