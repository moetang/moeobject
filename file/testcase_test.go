package file_test

import "github.com/spf13/afero"

var (
	fs afero.Fs
)

func prepare(f func()) {
	fs = afero.NewMemMapFs()
	if f != nil {
		f()
	}
}
