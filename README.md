# go-whosonfirst-fetch

Tools for fetching Who's On First records and their ancestors.

## Important

Work in progress. Documentation to follow.

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
$> make cli
go build -mod vendor -o bin/fetch cmd/fetch/main.go
```

### fetch

```
> ./bin/fetch -h
Usage of ./bin/fetch:
  -belongs-to value
    	One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.
  -reader-uri string
    	... (default "whosonfirst-data://")
  -retries int
    	... (default 3)
  -writer-uri string
    	 (default "null://")
```

## See also

* https://github.com/whosonfirst/go-reader
* https://github.com/whosonfirst/go-reader-whosonfirst-data
* https://github.com/whosonfirst/go-reader-http
* https://github.com/whosonfirst/go-writer