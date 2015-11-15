package cmd

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
)

var cmdZonesDetails = cli.Command{
	Name:      "details",
	Usage:     "print zone details",
	ArgsUsage: "<zone-id>",
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Usage error: zone id is required to print details.")
		}

		zone, err := client(c).Zones.Details(context.Background(), c.Args().First())
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
