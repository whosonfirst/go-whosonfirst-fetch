package fetch

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"github.com/whosonfirst/go-whosonfirst-uri"
	_ "io"
	_ "log"
	_ "os"
	"strings"
	"sync"
)

type Fetcher struct {
	reader     reader.Reader
	writer     writer.Writer
	processing *sync.Map
	processed  *sync.Map
	throttle   chan bool
	Force      bool
	Logger     *log.WOFLogger
}

func NewFetcher(rdr reader.Reader, wr writer.Writer) (*Fetcher, error) {

	logger := log.SimpleWOFLogger()

	processing := new(sync.Map)
	processed := new(sync.Map)

	max_fetch := 10
	throttle := make(chan bool, max_fetch)

	for i := 0; i < max_fetch; i++ {
		throttle <- true
	}

	f := Fetcher{
		reader:     rdr,
		writer:     wr,
		Force:      false,
		Logger:     logger,
		processing: processing,
		processed:  processed,
		throttle:   throttle,
	}

	return &f, nil
}

func (f *Fetcher) FetchIDs(ids []int64, belongs_to ...string) error {

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

func (f *Fetcher) FetchID(id int64, belongs_to ...string) error {

	_, ok := f.processed.Load(id)

	if ok {
		f.Logger.Status("%d has already been processed, skipping", id)
		return nil
	}

	_, ok = f.processing.LoadOrStore(id, true)

	if ok {
		f.Logger.Status("%d is being processed, skipping", id)
		return nil
	}

	f.Logger.Status("waiting for throttle (%d)", id)

	<-f.throttle

	f.Logger.Status("processing (%d)", id)

	defer func() {
		f.throttle <- true
		f.processing.Delete(id)
	}()

	path, err := uri.Id2RelPath(id)

	if err != nil {
		return err
	}

	f.Logger.Debug("fetch %d from %s and write to %s", id, f.reader.URI(path), f.writer.URI(path))

	infile, read_err := f.reader.Read(path)

	if read_err != nil {
		return read_err
	}

	defer func() {
		infile.Close()
	}()

	outpath := f.writer.URI(path)

	write_err := f.writer.Write(path, infile)

	if write_err != nil {
		return write_err
	}

	f.processed.Store(id, true)

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

			hiers := whosonfirst.Hierarchies(ft)

			for _, h := range hiers {

				for pt, other_id := range h {

					possible := true

					for _, candidate_id := range ids {

						if other_id == candidate_id {
							possible = false
							break
						}
					}

					if possible == false {
						continue
					}

					pt = strings.Replace(pt, "_id", "", -1)

					for _, candidate_pt := range belongs_to {

						if pt == candidate_pt {
							ids = append(ids, other_id)
							break
						}
					}

				}
			}
		}

		if len(ids) > 0 {

			err = f.FetchIDs(ids)

			if err != nil {
				// return err
			}
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
		// pass
	}

	err := f.FetchID(id, fetch_belongsto...)

	if err != nil {
		err_ch <- err
	}
}
