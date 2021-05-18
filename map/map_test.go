package map_test

import (
	"github.com/pubgo/bench/cmap"
	"github.com/valyala/fastrand"
	"github.com/zeebo/wyhash"

	"strconv"
	"sync"
	"testing"
)

var m = make(map[interface{}]interface{})

func BenchmarkMap_set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(int(fastrand.Uint32()))] = i
	}
}

func BenchmarkMap_get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m[strconv.Itoa(int(fastrand.Uint32()))]
	}
}

var sm sync.Map

func BenchmarkSyncMap_set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sm.Store(strconv.Itoa(int(fastrand.Uint32())), i)
	}
}

func BenchmarkSyncMap_get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = sm.Load(strconv.Itoa(int(fastrand.Uint32())))
	}
}

var cm = cmap.New()

func BenchmarkCMap_set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cm.Set(strconv.Itoa(int(fastrand.Uint32())), i)
	}
}

func BenchmarkCMap_get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cm.Get(strconv.Itoa(int(fastrand.Uint32())))
	}
}

func BenchmarkCompare(b *testing.B) {
	// sizes := []int{
	// 	0, 1, 3, 4, 8, 9, 16, 17, 32,
	// 	33, 64, 65, 96, 97, 128, 129, 240, 241,
	// 	512, 1024, 100 * 1024,
	// }

	// for _, size := range sizes {
	// 	b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
	// 		b.SetBytes(int64(size))
	// 		var acc uint64
	// 		d := string(make([]byte, size))
	// 		b.ReportAllocs()
	// 		b.ResetTimer()

	// 		for i := 0; i < b.N; i++ {
	// 			acc = wyhash.HashString(d, 0)
	// 		}
	// 		runtime.KeepAlive(acc)
	// 	})
	// }
}

var rngPool sync.Pool

func Uint64() uint64 {
	v := rngPool.Get()
	if v == nil {
		var dd wyhash.RNG
		v = &dd
	}

	r := v.(*wyhash.RNG)
	x := r.Uint64()
	rngPool.Put(r)
	return x
}

func BenchmarkCompareFastrand(b *testing.B) {
	b.Run("rand", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Uint64()
		}
	})

	b.Run("rand_p", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Uint64()
			}
		})
	})

	b.Run("fastrand", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fastrand.Uint32()
		}
	})

	b.Run("fastrand_p", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = fastrand.Uint32()
			}
		})
	})
}
