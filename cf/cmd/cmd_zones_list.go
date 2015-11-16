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
			Usage: "if true only prints ids",
		},
		cli.BoolFlag{
			Name:  "domains",
			Usage: "(only with --list) if true only prints domains",
		},
	},
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
			if c.Bool("list") {
				if c.Bool("domains") {
					fmt.Println(zone.Name)
				} else {
					fmt.Println(zone.ID)
				}
			} else {
				table.Append([]string{
					zone.ID,
					zone.Name,
					yesOrNo(zone.Paused),
					zone.Status,
				})
			}
		}

		if !c.Bool("list") {
			table.Render()
		}
	},
}
