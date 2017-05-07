package util

func Uint16LEBytes(n uint16) []byte {
	return []byte{byte(n & 0x00FF), byte((n >> 8) & 0x00FF)}
}

func BytesLEUint16(data []byte) uint16 {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return uint16(data[0])
	default:
		return (uint16(data[1])<<8)&0xFFFF | uint16(data[0])&0xFFFF
	}
}
