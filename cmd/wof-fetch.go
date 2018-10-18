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
	"github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"io"
	"log"
	_ "path/filepath"
)

func main() {

	str_valid := bundle.ValidReadersString()

	desc := fmt.Sprintf("DSN strings MUST contain a 'reader=SOURCE' pair followed by any additional pairs required by that reader. Supported reader sources are: %s.", str_valid)

	var reader_flags flags.MultiDSNString
	flag.Var(&reader_flags, "reader", desc)

	var mode = flag.String("mode", "repo", "...")

	// var target = flag.String("target", "", "Where to write the data fetched. Currently on filesystem targets are supported.")
	var fetch_belongsto = flag.Bool("fetch-belongsto", false, "Fetch all the IDs that a given ID belongs to.")
	var force = flag.Bool("force", false, "Fetch IDs even if they are already present.")

	flag.Parse()

	r, err := bundle.NewMultiReaderFromFlags(reader_flags)

	if err != nil {
		log.Fatal(err)
	}

	// PLEASE MAKE ME A MULTIWRITER THINGY...
	
	// data := filepath.Join(path, "data")
	// wr, err := writer.NewFSWriter(data)

	wr, err := writer.NewNullWriter()

	if err != nil {
		log.Fatal(err)
	}

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

		err = fetcher.FetchID(wofid, *fetch_belongsto)

		log.Println("FETCH", wofid, err)
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
