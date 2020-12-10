package lib

import "strconv"

// ParseInt parses the string as a base-10 int
func ParseInt(text string) (int, error) {
	value, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(value), nil
}

// ParseInt64 parses the string as a base-10 int64
func ParseInt64(text string) (int64, error) {
	return strconv.ParseInt(text, 10, 64)
}
