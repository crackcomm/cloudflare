package cmd

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

var cmdZonesList = cli.Command{
	Name:  "list",
	Usage: "lists zones",
	Action: func(c *cli.Context) {
		zones, err := client(c).Zones.List(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"Name",
			"Paused",
			"Status",
		})
		for _, zone := range zones {
			table.Append([]string{
				zone.ID,
				zone.Name,
				yesOrNo(zone.Paused),
				zone.Status,
			})
		}
		table.Render()
	},
}
