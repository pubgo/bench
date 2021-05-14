package cmap

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var m = New()
	m.Set("hello", nil)
	fmt.Println(m.Get("hello") == nil)
	fmt.Println(m.GetSet("hello1", nil))
	fmt.Println(m.GetSet("hello1", nil))
	fmt.Println(m.GetSet("hello1", "ok"))
	fmt.Println(m.GetSet("hello1", "ok"))
}
