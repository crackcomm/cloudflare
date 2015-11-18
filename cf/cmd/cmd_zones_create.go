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
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "print",
			Usage: "print json result",
		},
	},
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Usage error: domain name is required to create a zone.")
		}

		domain := c.Args().First()
		zone, err := client(c).Zones.Create(context.Background(), domain)
		if err != nil {
			log.Fatalf("Error creating zone: %v", err)
		}

		if c.Bool("print") {
			body, err := json.MarshalIndent(zone, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%s", body)
		} else {
			log.Printf("Zone %s created", domain)
		}
	},
}
