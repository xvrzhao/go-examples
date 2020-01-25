package unsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

// ArraySliceDiff 演示了 slice 与 array 本质上的区别。
// slice 是一个 reflect.SliceHeader 结构，而 array 是一段连续的内存结构。
func ArraySliceDiff() {
	s := []int{1, 2, 3}
	fmt.Printf("%p != %p == %x\n", &s, &(s[0]), (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)

	arr := [3]int{1, 2, 3}
	fmt.Printf("%p == %p\n", &arr, &arr[0])
}
