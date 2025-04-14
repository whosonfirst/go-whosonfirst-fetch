package fetch

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type RunOptions struct {
	ReaderURI  string
	WriterURI  string
	Retries    int
	MaxClients int
	BelongsTo  []string
	IDs        []int64
	Verbose    bool
}

func DeriveRunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVarsWithFeedback(fs, "OFFLINE", false)

	if err != nil {
		return nil, fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	uris := fs.Args()
	ids := make([]int64, 0)

	for _, raw := range uris {

		id, _, err := uri.ParseURI(raw)

		if err != nil {
			return nil, fmt.Errorf("Unable to parse URI '%s', %w", raw, err)
		}

		ids = append(ids, id)
	}

	opts := &RunOptions{
		ReaderURI:  reader_uri,
		WriterURI:  writer_uri,
		Retries:    retries,
		MaxClients: max_clients,
		BelongsTo:  belongs_to,
		IDs:        ids,
		Verbose:    verbose,
	}

	return opts, nil
}
