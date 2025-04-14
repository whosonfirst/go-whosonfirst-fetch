// fetch is a command line tool to retrieve one or more Who's on First records and, optionally, their ancestors.
package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-reader-http"
	_ "github.com/whosonfirst/go-reader-whosonfirst-data"

	"github.com/whosonfirst/go-whosonfirst-fetch/v2/app/fetch"
)

func main() {

	ctx := context.Background()
	err := fetch.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to fetch records, %v", err)
	}
}
