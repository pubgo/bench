package gc

import (
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func gcPrint() {
	var s = time.Now()
	runtime.GC()
	fmt.Println(time.Since(s))
	memStatsPrint()
	time.Sleep(time.Second)
}

func memStatsPrint() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024, m.HeapReleased/1024)
}

func TestStringMap(t *testing.T) {
	var dt = make(map[string]string)
	for i := 0; i < 1000000000; i++ {

		if i%100000 == 0 {
			gcPrint()
			fmt.Println(i)
		}

		var k = strconv.Itoa(i)
		dt[k] = k
	}
}

func TestIntMap(t *testing.T) {
	t.Run("int_string", func(t *testing.T) {
		var dt = make(map[int]string)
		for i := 0; i < 1000000000; i++ {

			if i%100000 == 0 {
				gcPrint()
				fmt.Println(i)
			}

			dt[i] = strconv.Itoa(i)
		}
	})

	t.Run("string_int", func(t *testing.T) {
		var dt = make(map[string]int)
		for i := 0; i < 1000000000; i++ {

			if i%100000 == 0 {
				gcPrint()
				fmt.Println(i)
			}

			dt[strconv.Itoa(i)] = i
		}
	})

	t.Run("int_int", func(t *testing.T) {
		var dt = make(map[int]int)
		for i := 0; i < 1000000000; i++ {

			if i%100000 == 0 {
				gcPrint()
				fmt.Println(i)
			}

			dt[i] = i
		}
	})

	t.Run("int_ptr_int", func(t *testing.T) {
		var dt = make(map[int]*int)
		for i := 0; i < 1000000000; i++ {

			if i%100000 == 0 {
				gcPrint()
				fmt.Println(i)
			}

			dt[i] = &i
		}
	})
}

type cc struct {
	checker copyChecker
}

func (t *cc) check() {
	t.checker.check()
}

func TestName11(t *testing.T) {
	var ccc = cc{}
	ccc.check()
	ccc.check()

	var c1 = ccc
	c1.check()
}

// copyChecker holds back pointer to itself to detect object copying.
type copyChecker uintptr

func (c *copyChecker) check() {
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("sync.Cond is copied")
	}
}
