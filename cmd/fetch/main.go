package main

/*
go run -mod vendor cmd/fetch/main.go -writer-uri 'fs:///usr/local/data/sfomuseum-data-whosonfirst/data' -belongs-to locality -belongs-to region -belongs-to country 102550865
*/

import (
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

import (
	_ "github.com/whosonfirst/go-reader-http"
	_ "github.com/whosonfirst/go-reader-whosonfirst-data"	
)

import (
	"context"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-writer"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"	
	"flag"
	"log"
	"strconv"
)

func main() {

	reader_uri := flag.String("reader-uri", "whosonfirst-data://", "...")
	writer_uri := flag.String("writer-uri", "null://", "")
	retries := flag.Int("retries", 3, "...")
	
	var belongs_to flags.MultiString
	flag.Var(&belongs_to, "belongs-to", "One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.")
	
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
	
	fetcher, err := fetch.NewFetcher(ctx, r, wr, fetcher_opts)

	if err != nil {
		log.Fatal(err)
	}
		
	str_ids := flag.Args()
	ids := make([]int64, 0)

	for _, str_id := range str_ids {

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, id)
	}
	
	err = fetcher.FetchIDs(ctx, ids, belongs_to...)

	if err != nil {
		log.Fatal(err)
	}
	
}
