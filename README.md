# Golang CloudFlare® API v4 client

[![GoDoc](https://godoc.org/github.com/crackcomm/cloudflare?status.svg)](https://godoc.org/github.com/crackcomm/cloudflare)

Golang API Client for CloudFlare® API v4.

## Usage

```go
package main

import (
	"log"
	"time"

	"github.com/crackcomm/cloudflare"

	"golang.org/x/net/context"
)

func main() {
	client := cloudflare.New(&cloudflare.Options{
		Email: "example@email.com",
		Key:   "example-key",
	})

	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*30))

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
