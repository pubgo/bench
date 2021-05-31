package binary

import (
	"testing"

	"github.com/kelindar/binary"
	"github.com/pubgo/bench/encode"
	"github.com/pubgo/xerror"
)

func BenchmarkBinaryEncode(b *testing.B) {
	s1 := encode.GetStu()
	for i := 0; i < b.N; i++ {
		if _, err := binary.Marshal(s1); err != nil {
			panic(err)
		}
	}
}

func BenchmarkBinaryDecode(b *testing.B) {
	s1 := encode.GetStu()
	var dt, err = binary.Marshal(s1)
	xerror.Panic(err)

	b.ResetTimer()
	var val encode.Student

	for i := 0; i < b.N; i++ {
		_ = binary.Unmarshal(dt, &val)
	}
}
