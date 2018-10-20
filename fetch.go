package fetch

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"github.com/whosonfirst/go-whosonfirst-uri"
	_ "log"
	_ "os"
)

type Fetcher struct {
	reader reader.Reader
	writer writer.Writer
	Force  bool
	Logger *log.WOFLogger
}

func NewFetcher(rdr reader.Reader, wr writer.Writer) (*Fetcher, error) {

	logger := log.SimpleWOFLogger()

	f := Fetcher{
		reader: rdr,
		writer: wr,
		Force:  false,
		Logger: logger,
	}

	return &f, nil
}

func (f *Fetcher) FetchIDs(ids []int64, belongs_to []string) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done_ch := make(chan bool)
	err_ch := make(chan error)

	for _, id := range ids {
		go f.FetchIDWithContext(ctx, id, belongs_to, done_ch, err_ch)
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

func (f *Fetcher) FetchID(id int64, belongs_to []string) error {

	path, err := uri.Id2RelPath(id)

	if err != nil {
		return err
	}

	f.Logger.Debug("fetch %d from %s and write to %s", id, f.reader.URI(path), f.writer.URI(path))

	outpath := f.writer.URI(path)
	do_fetch := true

	// this doesn't really make sense in a (multi) reader.Reader context - we
	// might need to add an 'Exists()' method to the reader.Reader interface...
	// (20181018/thisisaaronland)

	/*
		if !f.Force {
			_, err := os.Stat(outpath)
			do_fetch = os.IsNotExist(err)
		}
	*/

	if do_fetch {

		infile, err := f.reader.Read(path)

		if err != nil {
			return err
		}

		err = f.writer.Write(path, infile)
	}

	count_belongs_to := len(belongs_to)

	if count_belongs_to > 0 {

		ft, err := feature.LoadWOFFeatureFromFile(outpath)

		if err != nil {
			return err
		}

		ids := make([]int64, 0)

		if count_belongs_to == 1 && belongs_to[0] == "all" {

			ids = whosonfirst.BelongsTo(ft)

		} else {

			// CHECK wof:hierarchy...
		}

		err = f.FetchIDs(ids, []string{})

		if err != nil {
			return err
		}

	}

	return nil
}

func (f *Fetcher) FetchIDWithContext(ctx context.Context, id int64, fetch_belongsto []string, done_ch chan bool, err_ch chan error) {

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
