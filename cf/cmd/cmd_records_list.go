// Package cmd implements cloudflare cli commands.
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
)

var cmdRecordsList = cli.Command{
	Name:      "list",
	Usage:     "lists zone records",
	ArgsUsage: "<zone-id>",
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Usage error: zone id is required to print its records.")
		}

		records, err := client(c).Records.List(context.Background(), c.Args().First())
		if err != nil {
			log.Fatal(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"Type",
			"Name",
			"Content",
			"Proxiable",
			"Proxied",
			"Locked",
			"TTL",
			"Created On",
			"Modified On",
		})
		for _, record := range records {
			table.Append([]string{
				record.ID,
				record.Type,
				record.Name,
				record.Content,
				yesOrNo(record.Proxiable),
				yesOrNo(record.Proxied),
				yesOrNo(record.Locked),
				fmt.Sprintf("%d", record.TTL),
				record.CreatedOn.Format("2006/01/02 15:04:05"),
				record.ModifiedOn.Format("2006/01/02 15:04:05"),
			})
		}
		table.Render()
	},
}
