# Golang CloudFlare® API v4 client

[![GoDoc](https://godoc.org/github.com/crackcomm/cloudflare?status.svg)](https://godoc.org/github.com/crackcomm/cloudflare) [![Circle CI](https://img.shields.io/circleci/project/crackcomm/cloudflare.svg)](https://circleci.com/gh/crackcomm/cloudflare)


Golang API Client for CloudFlare® API v4.

## Command Line Tool

```sh
$ go install github.com/crackcomm/cloudflare/cf
$ cf
NAME:
   cf - CloudFlare command line tool

USAGE:
   cf [global options] command [command options] [arguments...]

VERSION:
   1.0.1

COMMANDS:
   zones	zones management
   records	zone records management
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --email 		CloudFlare user email [$CLOUDFLARE_EMAIL]
   --key 		CloudFlare user key [$CLOUDFLARE_KEY]
   --help, -h		show help
   --version, -v	print the version
```

## Usage

```go
package main

import (
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/crackcomm/cloudflare"
)

func main() {
	client := cloudflare.New(&cloudflare.Options{
		Email: "example@email.com",
		Key:   "example-key",
	})

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*30)

	zones, err := client.Zones.List(ctx)
	if err != nil {
		log.Fatal(err)
	} else if len(zones) == 0 {
		log.Fatal("No zones were found")
	}

	records, err := client.Records.List(ctx, zones[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		log.Printf("%#v", record)
	}
}
```

## CloudFlare®

CloudFlare is a registered trademark of [CloudFlare, Inc](https://cloudflare.com).

## License

Apache 2.0 License.
