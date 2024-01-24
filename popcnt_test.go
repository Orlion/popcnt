package popcnt

import (
	"math/bits"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopcntQuad(t *testing.T) {
	for i := 0; i < 1000; i++ {
		nums := [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()}
		assert.Equal(t, PopcntQuad(nums), serialPopcntQuad(nums))
	}
}

func BenchmarkPopcntQuad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopcntQuad([4]uint64{1, 2, 3, 4})
	}
}

func BenchmarkSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serialPopcntQuad([4]uint64{1, 2, 3, 4})
	}
}

func serialPopcntQuad(nums [4]uint64) [4]uint64 {
	return [4]uint64{uint64(bits.OnesCount64(nums[0])), uint64(bits.OnesCount64(nums[1])), uint64(bits.OnesCount64(nums[2])), uint64(bits.OnesCount64(nums[3]))}
}
