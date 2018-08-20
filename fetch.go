package fetch

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"log"
	"os"
)

// please add a logger thingy

type Fetcher struct {
	reader reader.Reader
	writer writer.Writer
	Force  bool
}

func NewFetcher(rdr reader.Reader, wr writer.Writer) (*Fetcher, error) {

	f := Fetcher{
		reader: rdr,
		writer: wr,
		Force:  false,
	}

	return &f, nil
}

func (f *Fetcher) FetchIDs(ids []int64, fetch_belongsto bool) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done_ch := make(chan bool)
	err_ch := make(chan error)

	for _, id := range ids {
		go f.FetchIDWithContext(ctx, id, fetch_belongsto, done_ch, err_ch)
	}

	remaining := len(ids)

	for remaining > 0 {

		select {
		case <-done_ch:
			remaining -= 1
		case err := <-err_ch:
			return err
		default:
			//
		}
	}

	return nil
}

func (f *Fetcher) FetchID(id int64, fetch_belongsto bool) error {

	path, err := uri.Id2RelPath(id)

	if err != nil {
		return err
	}

	log.Printf("fetch %d from %s and write to %s", id, f.reader.URI(path), f.writer.URI(path))

	outpath := f.writer.URI(path)
	do_fetch := true

	if !f.Force {
		_, err := os.Stat(outpath)
		do_fetch = os.IsNotExist(err)
	}

	log.Printf("do fetch for %s : %t\n", f.writer.URI(path), do_fetch)

	if do_fetch {

		infile, err := f.reader.Read(path)

		if err != nil {
			return err
		}

		err = f.writer.Write(path, infile)
	}

	if fetch_belongsto {

		ft, err := feature.LoadWOFFeatureFromFile(outpath)

		if err != nil {
			return err
		}

		ids := whosonfirst.BelongsTo(ft)

		err = f.FetchIDs(ids, false)

		if err != nil {
			return err
		}
	}

	return nil
}

func (f *Fetcher) FetchIDWithContext(ctx context.Context, id int64, fetch_belongsto bool, done_ch chan bool, err_ch chan error) {

	defer func() {
		done_ch <- true
	}()

	select {

	case <-ctx.Done():
		return
	default:

		err := f.FetchID(id, fetch_belongsto)

		if err != nil {
			err_ch <- err
		}
	}
}
