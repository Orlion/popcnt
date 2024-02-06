package popcnt

//go:noescape
func SimdPopcntQuad(nums [4]uint64) [4]uint64
