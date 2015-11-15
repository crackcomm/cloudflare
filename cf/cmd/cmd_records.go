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

var cmdRecordsDelete = cli.Command{
	Name:      "delete",
	Usage:     "deletes zone record",
	ArgsUsage: "<zone-id> <record-id>",
	Action:    func(c *cli.Context) {},
}
