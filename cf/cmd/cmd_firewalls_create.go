package cmd

import (
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/cloudflare"
)

var cmdFirewallsCreate = cli.Command{
	Name:  "create",
	Usage: "create a firewall rule",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "mode",
			Usage: "action mode [whitelist|block|challenge] (required)",
		},
		cli.StringFlag{
			Name:  "target",
			Usage: "target type [ip|ip_range|country] (required)",
		},
		cli.StringFlag{
			Name:  "value",
			Usage: "value is an ip, ip range or country code depending on the target value (required)",
		},
		cli.StringFlag{
			Name:  "notes",
			Usage: "optional note describing the change",
		},
	},
	Action: func(c *cli.Context) {
		zoneID, err := getZoneID(c)
		if err != nil {
			log.Fatal(err)
		}

		if c.String("mode") == "" {
			log.Fatal("Usage error: --mode flag is required.")
		}
		if c.String("target") == "" {
			log.Fatal("Usage error: --target flag is required.")
		}
		if c.String("value") == "" {
			log.Fatal("Usage error: --value flag is required.")
		}

		log.Printf("Creating firewall for zone: %s", zoneID)

		firewall := &cloudflare.Firewall{
			Mode: c.String("mode"),
			Configuration: &cloudflare.FirewallConfiguration{
				Target: c.String("target"),
				Value:  c.String("value"),
			},
			Notes: c.String("notes"),
		}

		fw, err := client(c).Firewalls.Create(context.Background(), zoneID, firewall)
		if err != nil {
			log.Fatalf("Error creating firewall: %v", err)
		}

		table := newFirewallsTable()
		table.add(fw)
		table.Render()
	},
}
