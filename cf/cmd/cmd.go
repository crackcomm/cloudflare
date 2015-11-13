// Package cmd implements cloudflare cli commands.
package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/crackcomm/cloudflare"
)

// New - Returns cloudflare cli commands.
func New() []cli.Command {
	return []cli.Command{
		cmdZones,
		cmdRecords,
	}
}

func client(c *cli.Context) *cloudflare.Client {
	return cloudflare.New(&cloudflare.Options{
		Key:   c.GlobalString("key"),
		Email: c.GlobalString("email"),
	})
}

func yesOrNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
