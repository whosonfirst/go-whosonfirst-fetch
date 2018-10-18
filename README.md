# go-whosonfirst-fetch

## Tools

### wof-fetch

Fetch, or refresh, a set of WOF records defined by an "index" from one or more "readers" and store them using one or more "writers". 

```
./bin/wof-fetch -h
Usage of ./bin/wof-fetch:
  -fetch-belongsto
    	Fetch all the IDs that a given ID belongs to.
  -mode string
    	The mode to use when indexing data. Valid modes are: directory, feature, feature-collection, files, geojson-ls, meta, path, repo, sqlite (default "repo")
  -reader value
    	One or more DSN strings representing a source to read data from. DSN strings MUST contain a 'reader=SOURCE' pair followed by any additional pairs required by that reader. Supported reader sources are: fs, github, http, mysql, repo, s3, sqlite.
  -retries int
    	The number of time to retry a failed fetch
  -writer value
    	One or more DSN strings representing a target to write data to. DSN strings MUST contain a 'writer=SOURCE' pair followed by any additional pairs required by that writer. Supported writer sources are: fs, null, repo, s3, sqlite, stdout.
```

An "index" is a valid `go-whosonfirst-index` thingy. A reader is a valid `go-whosonfirst-readwrite/reader` thingy. A writer is a valid `go-whosonfirst-readwrite/writer` thingy.

For example:

```
./bin/wof-fetch -writer 'writer=repo root=/usr/local/data/sfomuseum-data-whosonfirst' -reader 'reader=github repo=whosonfirst-data' -reader 'reader=github repo=whosonfirst-data-postalcode-us' -mode repo /usr/local/data/sfomuseum-data-whosonfirst/
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-index
* https://github.com/whosonfirst/go-whosonfirst-readwrite
