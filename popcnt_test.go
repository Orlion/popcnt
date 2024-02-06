package popcnt

import (
	"math/bits"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimdPopcntQuad(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1; i++ {
		nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
		assert.Equal(t, SimdPopcntQuad(nums), serialPopcntQuad(nums))
	}
}

func BenchmarkSimdPopcntQuad(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
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

func serialPopcntQuad(nums [4]uint64) [4]uint64 {
	return [4]uint64{uint64(bits.OnesCount64(nums[0])), uint64(bits.OnesCount64(nums[1])), uint64(bits.OnesCount64(nums[2])), uint64(bits.OnesCount64(nums[3]))}
}
