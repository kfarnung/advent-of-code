package lib

// Abs calculates the absolute value of a given signed integer
func Abs[T Signed](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

// Mod calculates the modulus value for signed integers (as opposed to the
// remainder)
// https://github.com/golang/go/issues/448
func Mod[T Signed](x, y T) T {
	return ((x % y) + y) % y
}
