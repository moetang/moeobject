package file

import (
	"errors"
	"io"
	"strconv"

	"github.com/moetang/moeobject/util"
)

type PhyFile struct {
	Id   uint16
	Type int
}

func NewPhyFile(w io.Writer, newPhyFileId uint16) (*PhyFile, error) {
	header := make([]byte, 16)
	header = header[:0]

	header = append(header, _PHY_FILE_HEADER...)
	id := util.Uint16LEBytes(newPhyFileId)
	header = append(header, id...)
	header = append(header, _PHY_FILE_TYPE_NORMAL...)
	header = append(header, _RESERVED_4_BYTES...)
	header = append(header, _PHY_FILE_FOOTER...)

	n, err := w.Write(header)
	if err != nil {
		return nil, err
	}
	if n != 16 {
		return nil, errors.New("write data length is not 16. only: " + strconv.Itoa(n))
	}
	return &PhyFile{
		Id:   newPhyFileId,
		Type: PHY_FILE_TYPE_NORMAL,
	}, nil
}

func LoadPhyFile(r io.ReadSeeker) (*PhyFile, error) {
	data := make([]byte, 16)
	_, err := io.ReadFull(r, data)
	if err != nil {
		return nil, err
	}
	if !util.CheckByteSliceEquals(data[0:4], _PHY_FILE_HEADER) {
		return nil, errors.New("PhyFile header not match.")
	}
	if !util.CheckByteSliceEquals(data[12:16], _PHY_FILE_FOOTER) {
		return nil, errors.New("PhyFile footer not match.")
	}
	id := util.BytesLEUint16(data[4:6])
	typeData := data[6:8]
	t := 0
	switch typeData[0] {
	case 0:
		switch typeData[1] {
		case 1:
			t = PHY_FILE_TYPE_NORMAL
		}
	}
	return &PhyFile{
		Id:   id,
		Type: t,
	}, nil
}
