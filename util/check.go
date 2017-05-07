package util

func CheckByteSliceEquals(src, dst []byte) bool {
	if len(src) != len(dst) {
		return false
	}
	for i := 0; i < len(src); i++ {
		if src[i] != dst[i] {
			return false
		}
	}
	return true
}
