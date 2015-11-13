package cmd

import "github.com/codegangsta/cli"

var cmdZones = cli.Command{
	Name:  "zones",
	Usage: "zones management",
	Subcommands: []cli.Command{
		cmdZonesCreate,
		cmdZonesList,
		cmdZonesDelete,
		cmdZonesDetails,
	},
}
