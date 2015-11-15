// Package cmd implements cloudflare cli commands.
package cmd

import "github.com/codegangsta/cli"

var cmdRecords = cli.Command{
	Name:  "records",
	Usage: "zone records management",
	Subcommands: []cli.Command{
		cmdRecordsList,
		cmdRecordsDelete,
	},
}
