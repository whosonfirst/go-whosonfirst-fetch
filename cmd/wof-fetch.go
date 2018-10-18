package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-index"
	"github.com/whosonfirst/go-whosonfirst-readwrite-bundle"
	_ "github.com/whosonfirst/go-whosonfirst-uri"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	str_valid := bundle.ValidReadersString()

	desc := fmt.Sprintf("DSN strings MUST contain a 'reader=SOURCE' pair followed by any additional pairs required by that reader. Supported reader sources are: %s.", str_valid)

	var reader_flags flags.MultiDSNString
	flag.Var(&reader_flags, "reader", desc)

	// var target = flag.String("target", "", "Where to write the data fetched. Currently on filesystem targets are supported.")
	var fetch_belongsto = flag.Bool("fetch-belongsto", true, "Fetch all the IDs that a given ID belongs to.")
	var force = flag.Bool("force", false, "Fetch IDs even if they are already present.")

	flag.Parse()

	r, err := bundle.NewMultiReaderFromFlags(reader_flags)

	if err != nil {
		log.Fatal(err)
	}

	// wr, err := bundle.NewMultiWriterFromFlags("...")

	var wr interface{}

	fetcher, err := fetch.NewFetcher(r, wr)

	if err != nil {
		log.Fatal(err)
	}

	fetcher.Force = *force

	cb := func(fh io.Reader, ctx context.Context, args ...interface{}) error {

		f, err := feature.LoadFeatureFromReader(fh)

		if err != nil {
			return err
		}

		wofid := whosonfirst.Id(f)

		return fetcher.FetchID(wofid, *fetch_belongsto)
	}

	i, err := index.NewIndexer(*mode, f)

	if err != nil {
		log.Fatal(err)
	}

	for _, path := range flag.Args() {

		err := i.IndexPath(path)

		if err != nil {
			log.Fatal(err)
		}
	}

}
