package lib

// AbsInt calculates the absolute value of a given int
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// ModInt64 calculates the modulus value (as opposed to the remainder)
// https://github.com/golang/go/issues/448
func ModInt64(x, y int64) int64 {
	return ((x % y) + y) % y
}
