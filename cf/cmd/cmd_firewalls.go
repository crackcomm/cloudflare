package cmd

import (
	"github.com/codegangsta/cli"
)

var cmdFirewalls = cli.Command{
	Name:  "firewall",
	Usage: "firewall management",
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
	Subcommands: []cli.Command{
		cmdFirewallsCreate,
		cmdFirewallsList,
		cmdFirewallsDelete,
	},
}
