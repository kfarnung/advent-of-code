package lib

// StringSliceToInt64 converts the slice of strings to a slice of int64 values.
func StringSliceToInt64(input []string) ([]int64, error) {
	result := make([]int64, len(input))
	for i, val := range input {
		num, err := ParseInt64(val)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}

// StringSliceToInt converts the slice of strings to a slice of int values.
func StringSliceToInt(input []string) ([]int, error) {
	result := make([]int, len(input))
	for i, val := range input {
		num, err := ParseInt(val)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}
