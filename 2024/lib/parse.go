package lib

import (
	"strconv"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// ParseInt64 parses the string as a base-10 int64
func ParseInt[T constraints.Signed](text string) (T, error) {
	size := unsafe.Sizeof(T(0))
	value, err := strconv.ParseInt(text, 10, int(size*8))
	if err != nil {
		return 0, err
	}

	return T(value), nil
}
