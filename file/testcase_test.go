package file_test

import (
	"io"
	"testing"

	"github.com/spf13/afero"
)

var (
	fs afero.Fs
)

func prepare(f func()) {
	fs = afero.NewMemMapFs()
	if f != nil {
		f()
	}
}

func errAssert(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func errAssertMsg(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatal(msg, "raw error:", err)
	}
}

func readAll(r io.Reader) ([]byte, error) {
	return afero.ReadAll(r)
}
