package fetch

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-reader"
	_ "github.com/whosonfirst/go-reader-whosonfirst-data"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"github.com/whosonfirst/go-writer"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFetch(t *testing.T) {

	ctx := context.Background()

	tmpdir, err := ioutil.TempDir("", "fetch")

	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println("TMPDIR", tmpdir)

	defer os.RemoveAll(tmpdir)

	r, err := reader.NewReader(ctx, "whosonfirst-data://")

	if err != nil {
		t.Fatal(err)
	}

	wr_uri := fmt.Sprintf("fs://%s", tmpdir)
	wr, err := writer.NewWriter(ctx, wr_uri)

	if err != nil {
		t.Fatal(err)
	}

	fetcher_opts, err := DefaultOptions()

	if err != nil {
		t.Fatal(err)
	}

	fetcher, err := NewFetcher(ctx, r, wr, fetcher_opts)

	if err != nil {
		t.Fatal(err)
	}

	ids := []int64{1360695651}
	belongs_to := []string{"all"}

	err = fetcher.FetchIDs(ctx, ids, belongs_to...)

	if err != nil {
		t.Fatal(err)
	}

	to_verify := []int64{
		1360695651,
		101756499,
		85633111,
		1377694369,
		102191581,
		102063913,
		85682553,
	}

	for _, id := range to_verify {

		rel_path, err := uri.Id2RelPath(id)

		if err != nil {
			t.Fatal(err)
		}

		abs_path := filepath.Join(tmpdir, rel_path)

		_, err = os.Stat(abs_path)

		if err != nil {
			t.Fatal(err)
		}
	}
}
