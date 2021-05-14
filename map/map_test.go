package map_test

import (
	"github.com/pubgo/bench/cmap"
	"github.com/valyala/fastrand"

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
