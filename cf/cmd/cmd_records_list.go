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
	ArgsUsage: "[zone]",
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Zone ID is required to print its records.")
		}

		zones, err := client(c).Records.List(context.Background(), c.Args().First())
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
		for _, zone := range zones {
			table.Append([]string{
				zone.ID,
				zone.Type,
				zone.Name,
				zone.Content,
				yesOrNo(zone.Proxiable),
				yesOrNo(zone.Proxied),
				yesOrNo(zone.Locked),
				fmt.Sprintf("%d", zone.TTL),
				zone.CreatedOn.Format("2006/01/02 15:04:05"),
				zone.ModifiedOn.Format("2006/01/02 15:04:05"),
			})
		}
		table.Render()
	},
}
