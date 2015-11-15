package cmd

import (
	"log"

	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var cmdZonesDelete = cli.Command{
	Name:      "delete",
	Usage:     "delete zone",
	ArgsUsage: "<zone-id>",
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Usage error: zone id is required to delete one.")
		}

		id := c.Args().First()
		err := client(c).Zones.Delete(context.Background(), id)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Deleted %s", id)
	},
}
