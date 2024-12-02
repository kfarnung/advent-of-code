package lib

import (
	"cmp"
	"sort"
)

// StringSliceToInt64 converts the slice of strings to a slice of int64 values.
func StringSliceToInt64(input []string) ([]int64, error) {
	result := make([]int64, len(input))
	for i, val := range input {
		num, err := ParseInt[int64](val)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}

// StringSliceToInt32 converts the slice of strings to a slice of int values.
func StringSliceToInt32(input []string) ([]int32, error) {
	result := make([]int32, len(input))
	for i, val := range input {
		num, err := ParseInt[int32](val)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}

func SortSliceAscending[T cmp.Ordered](input []T) {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
}
