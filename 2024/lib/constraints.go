package lib

// Signed is the set of signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
