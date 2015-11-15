// Package cmd implements cloudflare cli commands.
package cmd

import (
	"log"
	"os"

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
	opts := &cloudflare.Options{
		Key:   c.GlobalString("key"),
		Email: c.GlobalString("email"),
	}
	if opts.Key == "" || opts.Email == "" {
		log.Println("You have to provide Cloudflare Email and API key.")
		log.Println("Use CLOUDFLARE_EMAIL and CLOUDFLARE_KEY environment variables.")
		log.Println("Or alternatively provide them in -email and -key flags in each call.")
		os.Exit(255)
	}
	return cloudflare.New(opts)
}

func yesOrNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
