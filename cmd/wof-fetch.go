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
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-readwrite-bundle"
	"io"
	golog "log"
	"os"
	"strings"
)

func main() {

	valid_readers := bundle.ValidReadersString()
	valid_writers := bundle.ValidWritersString()

	desc_readers := fmt.Sprintf("One or more DSN strings representing a source to read data from. DSN strings MUST contain a 'reader=SOURCE' pair followed by any additional pairs required by that reader. Supported reader sources are: %s.", valid_readers)

	desc_writers := fmt.Sprintf("One or more DSN strings representing a target to write data to. DSN strings MUST contain a 'writer=SOURCE' pair followed by any additional pairs required by that writer. Supported writer sources are: %s.", valid_writers)

	var reader_flags flags.MultiDSNString
	flag.Var(&reader_flags, "reader", desc_readers)

	var writer_flags flags.MultiDSNString
	flag.Var(&writer_flags, "writer", desc_writers)

	valid_modes := index.Modes()
	str_valid_modes := strings.Join(valid_modes, ", ")

	desc_mode := fmt.Sprintf("The mode to use when indexing data. Valid modes are: %s", str_valid_modes)
	var mode = flag.String("mode", "repo", desc_mode)

	var fetch_belongsto = flag.Bool("fetch-belongsto", false, "Fetch all the IDs that a given ID belongs to.")

	var retries = flag.Int("retries", 0, "The number of time to retry a failed fetch")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	stdout := io.Writer(os.Stdout)
	logger.AddLogger(stdout, "status")

	r, err := bundle.NewMultiReaderFromFlags(reader_flags)

	if err != nil {
		golog.Fatal(err)
	}

	wr, err := bundle.NewMultiWriterFromFlags(writer_flags)

	if err != nil {
		golog.Fatal(err)
	}

	fetcher, err := fetch.NewFetcher(r, wr)

	if err != nil {
		golog.Fatal(err)
	}

	fetcher.Logger = logger

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

		if err != nil {
			logger.Warning("Unable to fetch %d because '%v'", wofid, err)
			return err
		}

		logger.Info("Successfully fetched %d", wofid)
		return nil
	}

	i, err := index.NewIndexer(*mode, cb)

	if err != nil {
		golog.Fatal(err)
	}

	for _, path := range flag.Args() {

		err = i.IndexPath(path)

		if err != nil {
			golog.Fatal(err)
		}
	}

}
