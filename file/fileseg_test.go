package file_test

import (
	"testing"

	"github.com/moetang/moeobject/file"
)

func TestNewFileSeg(t *testing.T) {
	prepare(nil)

	f, err := fs.Create("fileseg1.dat")
	errAssert(t, err)

	fileseg, err := file.NewFileSeg(f, 1)
	errAssert(t, err)
	t.Log(fileseg)

	f.Close()

	f, err = fs.Open("fileseg1.dat")
	errAssert(t, err)
	data, err := readAll(f)
	errAssert(t, err)
	t.Log(data)

	f.Close()
}
