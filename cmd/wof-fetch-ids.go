package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	github_reader "github.com/whosonfirst/go-whosonfirst-readwrite-github/reader"
	http_reader "github.com/whosonfirst/go-whosonfirst-readwrite-http/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"log"
	"os"
	"strconv"
)

func main() {

	var sources flags.MultiDSNString
	flag.Var(&sources, "reader", "...")

	// var source = flag.String("source", "fs", "Valid options are: fs, github")
	// var dsn = flag.String("dsn", "https://data.whosonfirst.org", "...")

	var target = flag.String("target", "", "Where to write the data fetched. Currently on filesystem targets are supported.")
	var fetch_belongsto = flag.Bool("fetch-belongsto", true, "Fetch all the IDs that a given ID belongs to.")
	var force = flag.Bool("force", false, "Fetch IDs even if they are already present.")

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

	var readers = make([]reader.Reader, 0)

	for _, dsn := range sources {

		reader, ok := dsn["type"]

		if !ok {
			log.Fatal("Missing source for DSN")
		}

		switch reader {
		case "http":

			// please write github_reader.NewHTTPReaderFromString(dsn string)

			src, ok := dsn["source"]

			if !ok {
				log.Fatal("Missing HTTP source")
			}

			r, err := http_reader.NewHTTPReader(src)

			if err != nil {
				log.Fatal(err)
			}

			readers = append(readers, r)

		case "github":

			// please write github_reader.NewGitHubReaderFromString(dsn string)

			repo, ok := dsn["repo"]

			if !ok {
				log.Fatal("Missing GitHub repo")
			}

			branch, ok := dsn["branch"]

			if !ok {
				branch = "master"
			}

			r, err := github_reader.NewGitHubReader(repo, branch)

			if err != nil {
				log.Fatal(err)
			}

			readers = append(readers, r)

		default:
			log.Fatal("Invalid source")
		}
	}

	if len(readers) == 0 {
		log.Fatal("At least one valid reader is required")
	}

	rdr, err := reader.NewMultiReader(readers...)

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

	// please make this a MultiWriter... maybe?

	wr, err := writer.NewFSWriter(*target)

	if err != nil {
		log.Fatal(err)
	}

	f, err := fetch.NewFetcher(rdr, wr)

	if err != nil {
		log.Fatal(err)
	}

	f.Force = *force

	err = f.FetchIDs(ids, *fetch_belongsto)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
