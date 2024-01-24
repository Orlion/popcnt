package popcnt

import (
	"math/bits"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPopcntQuad(t *testing.T) {
	for i := 0; i < 1000; i++ {
		nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
		assert.Equal(t, PopcntQuad(nums), serialPopcntQuad(nums))
	}
}

func TestSimdPopcntQuad(t *testing.T) {
	for i := 0; i < 1; i++ {
		nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
		assert.Equal(t, SimdPopcntQuad(nums), serialPopcntQuad(nums))
	}
}

func BenchmarkPopcntQuad(b *testing.B) {
	nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
	for i := 0; i < b.N; i++ {
		PopcntQuad(nums)
	}
}

func BenchmarkSimdPopcntQuad(b *testing.B) {
	nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
	for i := 0; i < b.N; i++ {
		SimdPopcntQuad(nums)
	}
}

func BenchmarkSerial(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
	for i := 0; i < b.N; i++ {
		serialPopcntQuad(nums)
	}
}

func BenchmarkSerialMy(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
	for i := 0; i < b.N; i++ {
		serialMyPopcntQuad(nums)
	}
}

func serialPopcntQuad(nums [4]uint64) [4]uint64 {
	return [4]uint64{uint64(bits.OnesCount64(nums[0])), uint64(bits.OnesCount64(nums[1])), uint64(bits.OnesCount64(nums[2])), uint64(bits.OnesCount64(nums[3]))}
}

func serialMyPopcntQuad(nums [4]uint64) [4]uint64 {
	return [4]uint64{uint64(myOnesCount64(nums[0])), uint64(myOnesCount64(nums[1])), uint64(myOnesCount64(nums[2])), uint64(myOnesCount64(nums[3]))}
}

func myOnesCount64(x uint64) int {
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}
