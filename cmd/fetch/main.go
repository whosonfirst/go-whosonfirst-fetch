// fetch is a command line tool to retrieve one or more Who's on First records and, optionally, their ancestors.
package main

import (
	_ "github.com/whosonfirst/go-reader-http"
	_ "github.com/whosonfirst/go-reader-whosonfirst-data"
)

import (
	"context"
	"flag"
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"github.com/whosonfirst/go-writer"
	"log"
	"os"
)

func main() {

	reader_uri := flag.String("reader-uri", "whosonfirst-data://", "A valid whosonfirst/go-reader URI.")
	writer_uri := flag.String("writer-uri", "null://", "A valid whosonfirst/go-writer URI.")
	retries := flag.Int("retries", 3, "The maximum number of attempts to try fetching a record.")
	max_clients := flag.Int("max-clients", 10, "The maximum number of concurrent requests for multiple Who's On First records.")

	var belongs_to multi.MultiString
	flag.Var(&belongs_to, "belongs-to", "One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Fetch one or more Who's on First records and, optionally, their ancestors.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [options] [path1 path2 ... pathN]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nNotes:\n\n")
		fmt.Fprintf(os.Stderr, wordwrap.WrapString("pathN may be any valid Who's On First ID or URI that can be parsed by the go-whosonfirst-uri package.\n\n", 80))
	}

	flag.Parse()

	ctx := context.Background()

	r, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatal(err)
	}

	wr, err := writer.NewWriter(ctx, *writer_uri)

	if err != nil {
		log.Fatal(err)
	}

	fetcher_opts, err := fetch.DefaultOptions()

	if err != nil {
		log.Fatal(err)
	}

	fetcher_opts.Retries = *retries
	fetcher_opts.MaxClients = *max_clients

	fetcher, err := fetch.NewFetcher(ctx, r, wr, fetcher_opts)

	if err != nil {
		log.Fatal(err)
	}

	uris := flag.Args()
	ids := make([]int64, 0)

	for _, raw := range uris {

		id, _, err := uri.ParseURI(raw)

		if err != nil {
			log.Fatalf("Unable to parse URI '%s', %v", raw, err)
		}

		ids = append(ids, id)
	}

	fetched_ids, err := fetcher.FetchIDs(ctx, ids, belongs_to...)

	if err != nil {
		log.Fatalf("Failed to fetch IDs, %v", err)
	}

	for _, id := range fetched_ids {
		fmt.Println(id)
	}

}
