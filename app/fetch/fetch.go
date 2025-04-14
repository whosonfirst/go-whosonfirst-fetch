// fetch is a command line tool to retrieve one or more Who's on First records and, optionally, their ancestors.
package fetch

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-writer/v3"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := DeriveRunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	r, err := reader.NewReader(ctx, opts.ReaderURI)

	if err != nil {
		return fmt.Errorf("Failed to create reader, %w", err)
	}

	wr, err := writer.NewWriter(ctx, opts.WriterURI)

	if err != nil {
		return fmt.Errorf("Failed to create writer, %w", err)
	}

	fetcher_opts, err := fetch.DefaultOptions()

	if err != nil {
		return fmt.Errorf("Failed to create fetch options, %w", err)
	}

	fetcher_opts.Retries = opts.Retries
	fetcher_opts.MaxClients = opts.MaxClients

	fetcher, err := fetch.NewFetcher(ctx, r, wr, fetcher_opts)

	if err != nil {
		return fmt.Errorf("Failed to create fetcher, %w", err)
	}

	slog.Debug("Fetch items", "count", len(opts.IDs))

	fetched_ids, err := fetcher.FetchIDs(ctx, opts.IDs, opts.BelongsTo...)

	if err != nil {
		return fmt.Errorf("Failed to fetch IDs, %w", err)
	}

	for _, id := range fetched_ids {
		fmt.Println(id)
	}

	return nil
}
