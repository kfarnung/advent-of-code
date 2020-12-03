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
