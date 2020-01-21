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
// TODO: Translate to English.
func StringHeader() {
	s := "xavier"

	fmt.Printf("%p\n", &s)                                                // 0xc000090490，StringHeader 结构体地址
	fmt.Println(uintptr(unsafe.Pointer(&s)))                              // 824634311824 == 0xc000090490，StringHeader 结构体地址
	fmt.Printf("%p\n", &(*reflect.StringHeader)(unsafe.Pointer(&s)).Data) // 0xc000090490，StringHeader 结构体第一个字段的地址，也是 StringHeader 结构体地址
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)         // 18791723，底层字符串真实的内存地址
}

// String2Bytes 演示了使用 unsafe 的方法将字符串转为字节切片。
// 因为 字符串内部结构 (reflect.StringHeader) 和切片的内部结构 (reflect.SliceHeader) 存在不同，直接转会出现问题。
//
// TODO: Translate to English.
func String2Bytes() {
	s := "xavier"

	// 直接转，将导致 sliceHeader 缺少 Cap 字段，读取到的 cap 值是 Len 字段后的内存中的值，存在不确定性。
	b1 := *(*[]byte)(unsafe.Pointer(&s))
	fmt.Println(len(b1), cap(b1)) // 6 17740064

	// 提取出 stringHeader 的字段，来构造 sliceHeader，再转为 []byte。
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	b2 := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}))
	fmt.Println(len(b2), cap(b2)) // 6 6

	// 构造一个匿名结构体，继承 string 的两个字段后再添加一个 Cap 属性，来模拟 sliceHeader。
	b3 := *(*[]byte)(unsafe.Pointer(&struct {
		string // 包含 Data 和 Len
		Cap    int
	}{s, len(s)}))
	fmt.Println(len(b3), cap(b3)) // 6 6
}

// ReadOnlyBytes 演示了只读的 bytes 切片。
// 通过字面量初始化的字符串，编译时会将内存设为只读，即使转换为 bytes 类型也不可
// 操控这部分内存，否则会抛出致命错误，无法通过 recover 捕获。
func ReadOnlyBytes() {
	s := "xavier"
	b := *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
	b[0] = 0x61 // throw fatal error
	fmt.Println(s, b)
}
