package file_test

import (
	"errors"
	"testing"

	"github.com/moetang/moeobject/file"
)

func TestLayerType(t *testing.T) {
	prepare(func() {
		f, err := fs.Create("phyfile.dat")
		if err != nil {
			t.Fatal(err)
		}
		n, err := f.Write([]byte{0x76, 0x3E, 0xB1, 0x3F})
		if err != nil {
			t.Fatal(err)
		}
		if n != 4 {
			t.Fatal(errors.New("write length is not 4."))
		}
		f.Close()

		f, err = fs.Create("fileseg.dat")
		if err != nil {
			t.Fatal(err)
		}
		n, err = f.Write([]byte{0xF1, 0x1E, 0xB1, 0x0C})
		if err != nil {
			t.Fatal(err)
		}
		if n != 4 {
			t.Fatal(errors.New("write length is not 4."))
		}
		f.Close()

		f, err = fs.Create("fileitem.dat")
		if err != nil {
			t.Fatal(err)
		}
		n, err = f.Write([]byte{0xD3, 0x14, 0xEF, 0x1A})
		if err != nil {
			t.Fatal(err)
		}
		if n != 4 {
			t.Fatal(errors.New("write length is not 4."))
		}
		f.Close()
	})

	f, err := fs.Open("phyfile.dat")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if file.GetLayerType(f) != file.LAYER_PHY_FILE {
			t.Fatal("PhyFile layer analyse error.", "cnt:", i)
		}
	}
	f.Close()

	f, err = fs.Open("fileseg.dat")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if file.GetLayerType(f) != file.LAYER_FILE_SEG {
			t.Fatal("FileSeg layer analyse error.", "cnt:", i)
		}
	}
	f.Close()

	f, err = fs.Open("fileitem.dat")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if file.GetLayerType(f) != file.LAYER_FILE_ITEM {
			t.Fatal("FileItem layer analyse error.", "cnt:", i)
		}
	}
	f.Close()
}
