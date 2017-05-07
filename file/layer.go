package file

import (
	"io"

	"github.com/moetang/moeobject/util"
)

func GetLayerType(r io.ReadSeeker) LayerType {
	data := make([]byte, 4)
	n, err := io.ReadFull(r, data)
	var readCnt int64 = int64(n)
	if err != nil {
		if readCnt > 0 {
			r.Seek(-readCnt, io.SeekCurrent)
		}
		return LAYER_ERROR
	}

	_, err = r.Seek(-readCnt, io.SeekCurrent)
	if err != nil {
		return LAYER_ERROR
	}
	if util.CheckByteSliceEquals(data, _PHY_FILE_HEADER) {
		return LAYER_PHY_FILE
	}
	if util.CheckByteSliceEquals(data, _FILE_SEG_HEADER) {
		return LAYER_FILE_SEG
	}
	if util.CheckByteSliceEquals(data, _FILE_ITEM_HEADER) {
		return LAYER_FILE_ITEM
	}
	return LAYER_UNKNOWN_TYPE
}
