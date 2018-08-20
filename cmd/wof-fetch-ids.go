package main

import (
	"flag"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	github_reader "github.com/whosonfirst/go-whosonfirst-readwrite-github/reader"
	http_reader "github.com/whosonfirst/go-whosonfirst-readwrite-http/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var source = flag.String("source", "fs", "Valid options are: fs, github")
	var dsn = flag.String("dsn", "https://data.whosonfirst.org", "...")
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

	// please make this a MultiReader...

	var rdr reader.Reader

	switch *source {

	case "http":

		// please write github_reader.NewHTTPReaderFromString(dsn string)

		r, err := http_reader.NewHTTPReader(*dsn)

		if err != nil {
			log.Fatal(err)
		}

		rdr = r

	case "github":

		// please write github_reader.NewGitHubReaderFromString(dsn string)

		*dsn = strings.Trim(*dsn, " ")
		parts := strings.Split(*dsn, "=")

		if len(parts) != 2 {
			log.Fatal("Invalid DSN")
		}

		if parts[0] != "repo" {
			log.Fatal("Invalid DSN")
		}

		repo := parts[1]

		r, err := github_reader.NewGitHubReader(repo, "master")

		if err != nil {
			log.Fatal(err)
		}

		rdr = r

	default:
		log.Fatal("Invalid source")
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
