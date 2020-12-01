package lib

import "strconv"

// StringSliceToInt64 converts the slice of strings to a slice of int64 values.
func StringSliceToInt64(input []string) ([]int64, error) {
	result := make([]int64, len(input))
	for i, val := range input {
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}
