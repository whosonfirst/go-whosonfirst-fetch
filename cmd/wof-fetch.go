package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-readwrite-bundle"
	"io"
	"log"
	_ "os"
)

func main() {

	str_valid := bundle.ValidReadersString()

	desc := fmt.Sprintf("DSN strings MUST contain a 'reader=SOURCE' pair followed by any additional pairs required by that reader. Supported reader sources are: %s.", str_valid)

	var reader_flags flags.MultiDSNString
	flag.Var(&reader_flags, "reader", desc)

	var writer_flags flags.MultiDSNString
	flag.Var(&writer_flags, "writer", "...")
	
	var mode = flag.String("mode", "repo", "...")

	// var target = flag.String("target", "", "Where to write the data fetched. Currently on filesystem targets are supported.")
	// var force = flag.Bool("force", false, "Fetch IDs even if they are already present.")

	var fetch_belongsto = flag.Bool("fetch-belongsto", false, "Fetch all the IDs that a given ID belongs to.")

	var retries = flag.Int("retries", 0, "The number of time to retry a failed fetch")

	flag.Parse()

	r, err := bundle.NewMultiReaderFromFlags(reader_flags)

	if err != nil {
		log.Fatal(err)
	}
	
	wr, err := bundle.NewMultiWriterFromFlags(writer_flags)

	if err != nil {
		log.Fatal(err)
	}

	fetcher, err := fetch.NewFetcher(r, wr)

	if err != nil {
		log.Fatal(err)
	}

	cb := func(fh io.Reader, ctx context.Context, args ...interface{}) error {

		f, err := feature.LoadFeatureFromReader(fh)

		if err != nil {
			return err
		}

		wofid := whosonfirst.Id(f)

		attempts := *retries + 1

		for attempts > 0 {

			err = fetcher.FetchID(wofid, *fetch_belongsto)
			attempts = attempts - 1

			if err == nil {
				break
			}
		}

		return err
	}

	i, err := index.NewIndexer(*mode, cb)

	if err != nil {
		log.Fatal(err)
	}

	for _, path := range flag.Args() {

		err = i.IndexPath(path)

		if err != nil {
			log.Fatal(err)
		}
	}

}
