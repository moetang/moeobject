package file_test

import (
	"io"
	"testing"

	"github.com/moetang/moeobject/file"
)

func TestNewPhyFile(t *testing.T) {
	prepare(nil)

	f, err := fs.Create("phyfile1.dat")
	if err != nil {
		t.Fatal(err)
	}

	phy, err := file.NewPhyFile(f, 1)
	if err != nil {
		t.Fatal(err)
	}

	if phy.Id != 1 {
		t.Fatal("PhyFile is not equals 1")
	}

	f.Close()

	f, err = fs.Open("phyfile1.dat")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]byte, 16)
	_, err = io.ReadFull(f, data)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{
		0x76, 0x3E, 0xB1, 0x3F,
		0x01, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00,
		0x52, 0x1F, 0x15, 0xEA,
	}

	for i := 0; i < len(data); i++ {
		if expected[i] != data[i] {
			t.Fatal("result does not match expected. result=", data, "expected=", expected)
		}
	}

	f.Close()
}

func TestLoadPhyFile(t *testing.T) {
	prepare(func() {
		f, err := fs.Create("phyfile1.dat")
		if err != nil {
			t.Fatal(err)
		}

		phy, err := file.NewPhyFile(f, 1)
		if err != nil {
			t.Fatal(err)
		}

		if phy.Id != 1 {
			t.Fatal("PhyFile is not equals 1")
		}

		f.Close()
	})

	f, err := fs.Open("phyfile1.dat")
	if err != nil {
		t.Fatal(err)
	}

	phy, err := file.LoadPhyFile(f)
	if err != nil {
		t.Fatal(err)
	}
	if phy == nil {
		t.Fatal("phy is nil.")
	}

	if phy.Id != 1 {
		t.Fatal("phy id not loaded.")
	}
	if phy.Type != file.PHY_FILE_TYPE_NORMAL {
		t.Fatal("phy type not correct.")
	}
}
