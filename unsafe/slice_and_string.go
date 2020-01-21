package unsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

// StringHeader 演示了 string 的内部数据结构。
// string 内部的数据结构为 reflect.StringHeader：
//
//   type StringHeader struct {
//   	// 底层字符串真实的内存地址
//   	Data uintptr
//   	// 字符串长度
//   	Len  int
//   }
//
// 直接打印字符串变量的地址其实是打印的 StringHeader 结构体的地址，字符串内存的真实地址为 Data 字段的值。
//
// TODO: 翻译成英文。
func StringHeader() {
	s := "xavier"

	fmt.Printf("%p\n", &s) // 0xc000090490，StringHeader 结构体地址
	fmt.Println(uintptr(unsafe.Pointer(&s))) // 824634311824 == 0xc000090490，StringHeader 结构体地址
	fmt.Printf("%p\n", &(*reflect.StringHeader)(unsafe.Pointer(&s)).Data) // 0xc000090490，StringHeader 结构体第一个字段的地址，也是 StringHeader 结构体地址
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)).Data) // 18791723，底层字符串真实的内存地址
}
