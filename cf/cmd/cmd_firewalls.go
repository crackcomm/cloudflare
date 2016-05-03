package cmd

import (
	"github.com/codegangsta/cli"
)

var cmdFirewalls = cli.Command{
	Name:  "firewall",
	Usage: "firewall management",
	Subcommands: []cli.Command{
		cmdFirewallsCreate,
		cmdFirewallsList,
		cmdFirewallsDelete,
	},
}
