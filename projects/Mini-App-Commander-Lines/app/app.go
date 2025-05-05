package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Return the commander lines app ready to be execute
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Commander Line App"
	app.Usage = "Search IPs and server names on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br", //Default value if not pass host
		},
	}

	app.Commands = []cli.Command{
		{
			Name: 	"ip",
			Usage: 	"Search address IPs on internet",
			Flags: 	flags,
			Action: searchIps, // function with logic
		},
		{
			Name: 	"servers",
			Usage: 	"Search servers name on internet",
			Flags: 	flags,
			Action: searchServers, // function with logic
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host) // name server
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}