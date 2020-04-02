package murmur3

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const (
	MAXN = 1073741823
	MINN = -1073741824
	CHARSET = "qwertyuioplkjhgfdsazxcvbnm0123456789"
)

func generatorSeed() int64 {
	src := rand.NewSource(time.Now().UnixNano())
	randN := rand.New(src)
	return int64(randN.Int31n(MAXN-MINN) + MINN)
}

func BenchmarkHashString128(b *testing.B)  {
	buf := make([]byte, 8192)
	var seed int64

	for length := 1; length <= cap(buf); length *= 2 {
		seed = generatorSeed()
		b.Run(strconv.Itoa(length), func(b *testing.B) {
			buf = buf[:length]
			b.SetBytes(int64(length))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				HashString(seed, string(buf))
			}
		})
	}
}

