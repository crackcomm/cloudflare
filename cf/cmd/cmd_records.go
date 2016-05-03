package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var cmdRecords = cli.Command{
	Name:  "records",
	Usage: "zone records management",
	Subcommands: []cli.Command{
		cmdRecordsCreate,
		cmdRecordsList,
		cmdRecordsDetails,
		cmdRecordsDelete,
	},
}

func getZoneID(c *cli.Context) (zoneID string, err error) {
	zoneID = c.String("zone")
	if zoneID != "" {
		return
	} else if c.String("domain") == "" {
		return "", fmt.Errorf("Usage error: --zone or --domain flag is required.")
	}
	if domain := c.String("domain"); domain != "" {
		zones, err := client(c).Zones.List(context.Background())
		if err != nil {
			return "", err
		}

		for _, zone := range zones {
			if zone.Name == domain {
				zoneID = zone.ID
				break
			}
		}

		if zoneID == "" {
			return "", fmt.Errorf("Domain %q was not found", domain)
		}
	}
	return
}
