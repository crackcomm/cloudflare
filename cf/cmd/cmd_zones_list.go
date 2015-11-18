package cmd

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

var cmdZonesList = cli.Command{
	Name:  "list",
	Usage: "lists zones",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "list",
			Usage: "print list instead of table",
		},
		cli.BoolFlag{
			Name:  "domains",
			Usage: "includes domains in list",
		},
		cli.BoolFlag{
			Name:  "domains-only",
			Usage: "prints only domains in list",
		},
	},
	Action: func(c *cli.Context) {
		zones, err := client(c).Zones.List(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if c.Bool("domains") && !c.Bool("list") {
			log.Fatal("Usage error: --domains can be only used with --list.")
		}

		if c.Bool("list") {
			for _, zone := range zones {
				if c.Bool("domains-only") || c.Bool("domains") {
					fmt.Println(zone.Name)
				}
				if !c.Bool("domains-only") {
					fmt.Println(zone.ID)
				}
			}
			return
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
