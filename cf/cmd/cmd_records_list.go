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
	Name:  "list",
	Usage: "lists zone records",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "list",
			Usage: "print list instead of table",
		},
	},
	Action: func(c *cli.Context) {
		zoneID, err := getZoneID(c)
		if err != nil {
			return
		}

		records, err := client(c).Records.List(context.Background(), zoneID)
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
