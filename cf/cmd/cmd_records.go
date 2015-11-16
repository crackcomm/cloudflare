package cmd

import "github.com/codegangsta/cli"

var cmdRecords = cli.Command{
	Name:  "records",
	Usage: "zone records management",
	Subcommands: []cli.Command{
		cmdRecordsCreate,
		cmdRecordsList,
		cmdRecordsDelete,
	},
}
