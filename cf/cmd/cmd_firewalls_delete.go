package cmd

import (
	"log"

	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var cmdFirewallsDelete = cli.Command{
	Name:      "delete",
	Usage:     "deletes firewall rule",
	ArgsUsage: "<rule-id> [<rule-id> ...]",
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

		if len(c.Args()) == 0 {
			log.Fatal("Usage error: at least one rule id is required.")
		}

		for _, id := range c.Args() {
			err := client(c).Firewalls.Delete(context.Background(), zoneID, id)
			if err != nil {
				log.Fatalf("Error deleting %q: %v", id, err)
			}
			log.Printf("Deleted firewall rule with id %q.", id)
		}
	},
}
