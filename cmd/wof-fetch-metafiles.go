package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-csv"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-whosonfirst-readwrite-fs/writer"
	"github.com/whosonfirst/go-whosonfirst-readwrite-http/reader"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	var source = flag.String("source", "https://data.whosonfirst.org", "...")
	var target = flag.String("target", "", "...")
	var fetch_hierarchy = flag.Bool("fetch-hierarchy", true, "...")
	var force = flag.Bool("false", false, "...")

	flag.Parse()

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

	ids := make([]int64, 0)

	for _, path := range flag.Args() {

		reader, err := csv.NewDictReaderFromPath(path)

		if err != nil {
			log.Fatal(err)
		}

		for {
			row, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			str_id := row["id"]

			id, err := strconv.ParseInt(str_id, 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			ids = append(ids, id)

			if len(ids) == 10 {

				err = f.FetchIDs(ids, *fetch_hierarchy)

				if err != nil {
					log.Fatal(err)
				}

				ids = make([]int64, 0)
			}
		}
	}

	if len(ids) > 0 {

		err = f.FetchIDs(ids, *fetch_hierarchy)

		if err != nil {
			log.Fatal(err)
		}
	}

	os.Exit(0)
}
