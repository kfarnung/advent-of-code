package lib

import (
	"strconv"
	"unsafe"
)

// ParseInt parses the string as a base-10 signed integer of type T
func ParseInt[T Signed](text string) (T, error) {
	size := unsafe.Sizeof(T(0))
	value, err := strconv.ParseInt(text, 10, int(size*8))
	if err != nil {
		return 0, err
	}

	return T(value), nil
}
