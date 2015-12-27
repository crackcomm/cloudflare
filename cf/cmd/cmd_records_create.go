package cmd

import (
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/cloudflare"
)

var cmdRecordsCreate = cli.Command{
	Name:  "create",
	Usage: "creates zone record",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "record name (required)",
		},
		cli.StringFlag{
			Name:  "content",
			Usage: "record content (required)",
		},
		cli.StringFlag{
			Name:  "type",
			Usage: "record type (required)",
		},
		cli.BoolFlag{
			Name:  "proxied",
			Usage: "should record be proxied",
		},
		cli.IntFlag{
			Name:  "ttl",
			Value: 3600,
			Usage: "record time to live",
		},
	},
	Action: func(c *cli.Context) {
		zoneID, err := getZoneID(c)
		if err != nil {
			return
		}
		if c.String("type") == "" {
			log.Fatal("Usage error: --type flag is required.")
		}
		if c.String("name") == "" {
			log.Fatal("Usage error: --name flag is required.")
		}
		if c.String("content") == "" {
			log.Fatal("Usage error: --content flag is required.")
		}

		log.Printf("Creating record in zone: %s", zoneID)

		record := &cloudflare.Record{
			Type:    c.String("type"),
			Name:    c.String("name"),
			Content: c.String("content"),
			TTL:     c.Int("ttl"),
			Proxied: c.Bool("proxied"),
			ZoneID:  zoneID,
		}

		if err := client(c).Records.Create(context.Background(), record); err != nil {
			log.Fatalf("Error creating record: %v", err)
		}

		table := newRecordsTable()
		table.add(record)
		table.Render()
	},
}
