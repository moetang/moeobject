package file

import (
	"errors"
	"io"

	"github.com/moetang/moeobject/util"
)

type FileSeg struct {
	Id               uint16
	FileFormatFirst  int
	FileFormatSecond int
	Flags            []byte
	Version          []byte

	VariableHeaderData []byte
}

func NewFileSeg(w io.Writer, newId uint16) (*FileSeg, error) {
	seg := new(FileSeg)
	headerLen := 1024 // currently only support 1K header
	var headerFixed []byte
	var headerVariable []byte = make([]byte, headerLen-24)

	headerFixed = append(headerFixed, _FILE_SEG_HEADER...)
	seg.FileFormatFirst = FILE_SEG_11_SIMPLE
	seg.FileFormatSecond = FILE_SEG_21_VARIABLE_LENGTH
	headerFixed = append(headerFixed, byte(FILE_SEG_11_SIMPLE), byte(FILE_SEG_21_VARIABLE_LENGTH))
	idData := util.Uint16LEBytes(newId)
	seg.Id = newId
	headerFixed = append(headerFixed, idData...)
	seg.Flags = _EMPTY_4_BYTES
	headerFixed = append(headerFixed, _EMPTY_4_BYTES...)
	headerLenData := util.Uint16LEBytes(uint16(headerLen / 8))
	headerFixed = append(headerFixed, headerLenData...)
	seg.Version = []byte{0x00, 0x01}
	headerFixed = append(headerFixed, 0x00, 0x01)
	headerFixed = append(headerFixed, _EMPTY_4_BYTES...) //TODO checksum
	headerFixed = append(headerFixed, _FILE_SEG_FOOTER...)

	// fill variable part
	seg.VariableHeaderData = headerVariable
	headerFixed = append(headerFixed, headerVariable...)

	n, err := w.Write(headerFixed)
	if n != headerLen {
		return nil, errors.New("file seg header length not match.")
	}
	if err != nil {
		return nil, err
	}

	return seg, nil
}

func LoadFileSeg(r io.ReadSeeker) (*FileSeg, error) {
	//TODO
	return nil, nil
}
