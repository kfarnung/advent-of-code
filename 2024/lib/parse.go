package lib

import "strconv"

// ParseInt32 parses the string as a base-10 int
func ParseInt32(text string) (int32, error) {
	value, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(value), nil
}

// ParseInt64 parses the string as a base-10 int64
func ParseInt64(text string) (int64, error) {
	return strconv.ParseInt(text, 10, 64)
}
