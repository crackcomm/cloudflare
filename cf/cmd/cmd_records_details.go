package cmd

import (
	"encoding/json"
	"log"

	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var cmdRecordsDetails = cli.Command{
	Name:      "details",
	Usage:     "prints record details",
	ArgsUsage: "<record-id>",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "domain",
			Usage: "domain name",
		},
		cli.StringFlag{
			Name:  "zone",
			Usage: "zone id",
		},
	},
	Action: func(c *cli.Context) {
		zoneID, err := getZoneID(c)
		if err != nil {
			log.Fatal(err)
		}

		id := c.Args().First()
		if id == "" {
			cli.ShowSubcommandHelp(c)
		}

		record, err := client(c).Records.Details(context.Background(), zoneID, id)
		if err != nil {
			log.Fatalf("Error getting details for id %q: %v", id, err)
		}

		body, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s", body)
	},
}
