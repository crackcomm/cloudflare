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

var cmdRecordsList = cli.Command{
	Name:      "list",
	Usage:     "lists zone records",
	ArgsUsage: "<zone-id>",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "list",
			Usage: "print list instead of table",
		},
		cli.StringFlag{
			Name:  "domain",
			Usage: "print list instead of table",
		},
	},
	Action: func(c *cli.Context) {
		zoneID := c.Args().First()
		if zoneID == "" && c.String("domain") == "" {
			log.Fatal("Usage error: zone id or --domain is required to print its records.")
		}

		cfclient := client(c)

		if domain := c.String("domain"); domain != "" {
			zones, err := client(c).Zones.List(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			for _, zone := range zones {
				if zone.Name == domain {
					zoneID = zone.ID
					break
				}
			}

			if zoneID == "" {
				log.Fatalf("Domain %q was not found", domain)
			}
		}

		records, err := cfclient.Records.List(context.Background(), zoneID)
		if err != nil {
			log.Fatal(err)
		}

		if c.Bool("list") {
			for _, record := range records {
				fmt.Println(record.ID)
			}
			return
		}

		table := newRecordsTable()
		for _, record := range records {
			table.add(record)
		}
		table.Render()
	},
}

type recordsTable struct {
	table *tablewriter.Table
}

func newRecordsTable() *recordsTable {
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
	return &recordsTable{
		table: table,
	}
}

func (table *recordsTable) Render() { table.table.Render() }

func (table *recordsTable) add(record *cloudflare.Record) {
	table.table.Append([]string{
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
