package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestMemory(t *testing.T) {
	s := "123"
	s_size := len(s)
	fmt.Println(s_size)
	s_memory := unsafe.Sizeof(s)
	s_memory2 := reflect.TypeOf(s).Size()
	fmt.Println(s_memory)
	fmt.Println(s_memory2)
	s_bytes := []byte(s)
	fmt.Println(s_bytes)
}
