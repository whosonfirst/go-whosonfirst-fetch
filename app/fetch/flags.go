package fetch

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/go-wordwrap"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var reader_uri string
var writer_uri string
var retries int
var max_clients int
var belongs_to multi.MultiString
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("fetch")

	fs.StringVar(&reader_uri, "reader-uri", "whosonfirst-data://", "A valid whosonfirst/go-reader URI.")
	fs.StringVar(&writer_uri, "writer-uri", "stdout://", "A valid whosonfirst/go-writer URI.")
	fs.IntVar(&retries, "retries", 3, "The maximum number of attempts to try fetching a record.")
	fs.IntVar(&max_clients, "max-clients", 10, "The maximum number of concurrent requests for multiple Who's On First records.")

	fs.Var(&belongs_to, "belongs-to", "One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Fetch one or more Who's on First records and, optionally, their ancestors.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [options] [path1 path2 ... pathN]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nNotes:\n\n")
		fmt.Fprintf(os.Stderr, wordwrap.WrapString("pathN may be any valid Who's On First ID or URI that can be parsed by the go-whosonfirst-uri package.\n\n", 80))
	}

	return fs
}
