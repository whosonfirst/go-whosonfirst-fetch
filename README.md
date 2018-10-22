# go-whosonfirst-fetch

## Tools

### wof-fetch

Fetch, or refresh, a set of WOF records defined by an "index" from one or more "readers" and store them using one or more "writers". 

```
./bin/wof-fetch -h
Usage of ./bin/wof-fetch:
  -belongs-to value
    	One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.
  -clients int
    	The number of time to retry a failed fetch. (default 10)
  -mode string
    	The mode to use when indexing data. Valid modes are: directory, feature, feature-collection, files, geojson-ls, meta, path, repo, sqlite (default "repo")
  -reader value
    	One or more DSN strings representing a source to read data from. DSN strings MUST contain a 'reader=SOURCE' pair followed by any additional pairs required by that reader. Supported reader sources are: fs, github, http, mysql, repo, s3, sqlite.
  -retries int
    	The number of time to retry a failed fetch.
  -timings
    	Display timings when fetching records.
  -writer value
    	One or more DSN strings representing a target to write data to. DSN strings MUST contain a 'writer=SOURCE' pair followed by any additional pairs required by that writer. Supported writer sources are: fs, null, repo, s3, sqlite, stdout.
```

An "index" is a valid `go-whosonfirst-index` thingy. A reader is a valid `go-whosonfirst-readwrite/reader` thingy. A writer is a valid `go-whosonfirst-readwrite/writer` thingy.

For example:

```
./bin/wof-fetch -writer 'writer=repo root=/usr/local/data/sfomuseum-data-whosonfirst' -reader 'reader=github repo=whosonfirst-data' -reader 'reader=github repo=whosonfirst-data-postalcode-us' -mode repo /usr/local/data/sfomuseum-data-whosonfirst/
```

Or:

```
./bin/wof-fetch -timings -belongs-to country -writer 'writer=repo root=/usr/local/data/sfomuseum-data-whosonfirst' -reader 'reader=repo root=/usr/local/data/whosonfirst-data' -reader 'reader=github repo=whosonfirst-data-postalcode-us' /usr/local/data/sfomuseum-data-whosonfirst
12:19:49.914853 [wof-fetch] STATUS Time to process 85632163: 5.76848ms
12:19:49.982293 [wof-fetch] STATUS 85632315 has already been processed, skipping
12:19:49.988463 [wof-fetch] STATUS Time to process 85632315: 58.509777ms
12:19:50.130251 [wof-fetch] STATUS 85632179 has already been processed, skipping
12:19:50.130457 [wof-fetch] STATUS Time to process 85632179: 93.202693ms
... and so on
```

#### Known-knowns

* It is possible to trigger network and/or GitHub rate-limits if you are processing a large number of "belongs-to" relations and reading from GitHub. I haven't sorted that out yet.

## See also

* https://github.com/whosonfirst/go-whosonfirst-index
* https://github.com/whosonfirst/go-whosonfirst-readwrite
* https://github.com/whosonfirst/go-whosonfirst-readwrite-bundle
