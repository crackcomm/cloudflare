package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"

	"github.com/crackcomm/cloudflare"
)

var cmdFirewallsList = cli.Command{
	Name:  "list",
	Usage: "lists firewall rules",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "list",
			Usage: "print list instead of table",
		},
	},
	Action: func(c *cli.Context) {
		zoneID, err := getZoneID(c)
		if err != nil {
			log.Fatal(err)
		}

		rules, err := client(c).Firewalls.List(context.Background(), zoneID)
		if err != nil {
			log.Fatal(err)
		}

		if c.Bool("list") {
			for _, rule := range rules {
				fmt.Println(rule.ID)
			}
			return
		}

		table := newFirewallsTable()
		for _, rule := range rules {
			table.add(rule)
		}
		table.Render()
	},
}

type firewallsTable struct {
	table *tablewriter.Table
}

func newFirewallsTable() *firewallsTable {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID",
		"Mode",
		"Target",
		"Value",
		"Notes",
		"Created On",
		"Modified On",
	})
	return &firewallsTable{
		table: table,
	}
}

func (table *firewallsTable) Render() { table.table.Render() }

func (table *firewallsTable) add(firewall *cloudflare.Firewall) {
	table.table.Append([]string{
		firewall.ID,
		firewall.Mode,
		firewall.Configuration.Target,
		firewall.Configuration.Value,
		firewall.Notes,
		firewall.CreatedOn.Format("2006/01/02 15:04:05"),
		firewall.ModifiedOn.Format("2006/01/02 15:04:05"),
	})
}
