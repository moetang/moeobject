package file

var (
	_RESERVED_4_BYTES = []byte{0x00, 0x00, 0x00, 0x00}
	_EMPTY_4_BYTES    = _RESERVED_4_BYTES

	_PHY_FILE_HEADER      = []byte{0x76, 0x3E, 0xB1, 0x3F}
	_PHY_FILE_FOOTER      = []byte{0x52, 0x1F, 0x15, 0xEA}
	_PHY_FILE_TYPE_NORMAL = []byte{0x00, 0x01}

	_FILE_SEG_HEADER = []byte{0xF1, 0x1E, 0xB1, 0x0C}
	_FILE_SEG_FOOTER = []byte{0x52, 0x1F, 0x15, 0xEB}

	_FILE_ITEM_HEADER = []byte{0xD3, 0x14, 0xEF, 0x1A}
)

const (
	PHY_FILE_TYPE_NORMAL = 1

	FILE_SEG_11_SIMPLE          = 1
	FILE_SEG_12_REPLICA         = 2
	FILE_SEG_13_EC              = 3
	FILE_SEG_21_VARIABLE_LENGTH = 1
	FILE_SEG_22_FIXED_LENGTH    = 2
)

type LayerType int

const (
	LAYER_PHY_FILE = iota
	LAYER_FILE_SEG
	LAYER_FILE_ITEM
	LAYER_UNKNOWN_TYPE
	LAYER_ERROR
)
