package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-whosonfirst-readwrite-fs/writer"
	"github.com/whosonfirst/go-whosonfirst-readwrite-http/reader"
	"log"
	"os"
	"strconv"
)

func main() {

	var source = flag.String("source", "https://data.whosonfirst.org", "...")
	var target = flag.String("target", "", "...")
	var fetch_hierarchy = flag.Bool("fetch-hierarchy", true, "...")
	var force = flag.Bool("force", false, "...")

	flag.Parse()

	ids := make([]int64, 0)

	for _, str_id := range flag.Args() {

		id, err := strconv.ParseInt(str_id, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, id)
	}

	if len(ids) == 0 {
		log.Fatal("Nothing to fetch!")
	}

	rdr, err := reader.NewHTTPReader(*source)

	if err != nil {
		log.Fatal(err)
	}

	if *target == "" {

		cwd, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		*target = cwd
	}

	wr, err := writer.NewFSWriter(*target)

	if err != nil {
		log.Fatal(err)
	}

	f, err := fetch.NewFetcher(rdr, wr)

	if err != nil {
		log.Fatal(err)
	}

	f.Force = *force

	err = f.FetchIDs(ids, *fetch_hierarchy)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
