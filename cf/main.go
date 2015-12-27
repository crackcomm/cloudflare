package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/cloudflare/cf/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	app := cli.NewApp()
	app.Name = "cf"
	app.HelpName = app.Name
	app.Usage = "CloudFlare command line tool"
	app.Version = "1.0.0"
	app.Commands = cmd.New()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "email",
			Usage:  "user email",
			EnvVar: "CLOUDFLARE_EMAIL",
		},
		cli.StringFlag{
			Name:   "key",
			Usage:  "user key",
			EnvVar: "CLOUDFLARE_KEY",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
