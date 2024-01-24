package popcnt

const m0 = 0x5555555555555555 // 01010101 ...
const m1 = 0x3333333333333333 // 00110011 ...
const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...
const m = 1<<7 - 1

func PopcntQuad(nums [4]uint64) [4]uint64 {
	// x = x>>1&(m0&m) + x&(m0&m)
	t1 := andQuad(shiftRight1Quad(nums), [4]uint64{m0, m0, m0, m0})
	t2 := andQuad(nums, [4]uint64{m0, m0, m0, m0})
	nums = addQuad(t1, t2)

	// x = x>>2&(m1&m) + x&(m1&m)
	t1 = andQuad(shiftRight2Quad(nums), [4]uint64{m1, m1, m1, m1})
	t2 = andQuad(nums, [4]uint64{m1, m1, m1, m1})
	nums = addQuad(t1, t2)

	// x = (x>>4 + x) & (m2 & m)
	nums = andQuad(addQuad(shiftRight4Quad(nums), nums), [4]uint64{m2, m2, m2, m2})

	// x += x >> 8
	nums = addQuad(nums, shiftRight8Quad(nums))

	// x += x >> 16
	nums = addQuad(nums, shiftRight16Quad(nums))

	// x += x >> 32
	nums = addQuad(nums, shiftRight32Quad(nums))

	// int(x) & (1<<7 - 1)
	return andQuad(nums, [4]uint64{m, m, m, m})
}

//go:noescape
func SimdPopcntQuad(nums [4]uint64) [4]uint64

//go:noescape
func shiftRight1Quad(nums [4]uint64) [4]uint64

//go:noescape
func shiftRight2Quad(nums [4]uint64) [4]uint64

//go:noescape
func shiftRight4Quad(nums [4]uint64) [4]uint64

//go:noescape
func shiftRight8Quad(nums [4]uint64) [4]uint64
func shiftRight16Quad(nums [4]uint64) [4]uint64

//go:noescape
func shiftRight32Quad(nums [4]uint64) [4]uint64

//go:noescape
func andQuad(a [4]uint64, b [4]uint64) [4]uint64

//go:noescape
func addQuad(a [4]uint64, b [4]uint64) [4]uint64
