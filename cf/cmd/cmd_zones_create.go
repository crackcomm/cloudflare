package cmd

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
)

var cmdZonesCreate = cli.Command{
	Name:      "create",
	Usage:     "creates zone",
	ArgsUsage: "<domain>",
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Domain name is required to create zone.")
		}

		domain := c.Args().First()
		zone, err := client(c).Zones.Create(context.Background(), domain)
		if err != nil {
			log.Fatal(err)
		}

		body, err := json.MarshalIndent(zone, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s", body)
	},
}
