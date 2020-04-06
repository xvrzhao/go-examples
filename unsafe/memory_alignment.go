package unsafe

import (
	"fmt"
	"unsafe"
)

// Note: All annotations are results of execution under 64-bit platform.

func Sizeof() {
	// int
	fmt.Println(
		unsafe.Sizeof(int(0)),     // 8
		unsafe.Sizeof(uint(0)),    // 8
		unsafe.Sizeof(int8(0)),    // 1
		unsafe.Sizeof(uint8(0)),   // 1
		unsafe.Sizeof(int16(0)),   // 2
		unsafe.Sizeof(uint16(0)),  // 2
		unsafe.Sizeof(int32(0)),   // 4
		unsafe.Sizeof(uint32(0)),  // 4
		unsafe.Sizeof(int64(0)),   // 8
		unsafe.Sizeof(uint64(0)),  // 8
		unsafe.Sizeof(uintptr(0)), // 8
	)

	// float
	fmt.Println(unsafe.Sizeof(float32(0)), unsafe.Sizeof(float64(0))) // 4 8

	// bool
	fmt.Println(unsafe.Sizeof(false), unsafe.Sizeof(true)) // 1 1

	// string is underlying type reflect.StringHeader
	str := "xvr"
	fmt.Println(
		unsafe.Sizeof(str),       // 16 uintptr + int
		unsafe.Sizeof("xvrzhao"), // 16
	)

	// slice is underlying type reflect.SliceHeader, though s is not allocated memory (s == nil)
	var s0 []int8
	var s1 []int16
	fmt.Println(
		unsafe.Sizeof(s0),                 // 24 uintptr + int + int
		unsafe.Sizeof(s1),                 // 24
		unsafe.Sizeof([]int8{}),           // 24
		unsafe.Sizeof([]int8{0, 1, 2, 3}), // 24
	)
}

// StructCompare 示例说明了：具有相同字段的结构体，不同的字段排序方式，所产生的内存占用不同，
// 原因是由于结构体字段和结构体整体的对齐方式。
//
// 知识点：
//   1. 64 位平台的机器字长（CPU 一次可读入的字节数）为 8 字节，32 位平台为 4 字节。
//   2. 一个数据类型在对齐时，要保证其起始地址为其对齐值的整数倍。
//   3. 结构体的字节长度为 最长字节的字段的对齐值(与结构体的对齐值相等，通过 unsafe.Alignof(结构体实例) 计算) 的最小整数倍。
//
// 参考链接：
//   - https://www.jianshu.com/p/49f7e6f56568
//   - https://www.bilibili.com/video/BV1iZ4y1j7TT
//   - https://github.com/talk-go/night/issues/588
func StructCompare() {
	type S1 struct {
		a int8  // 1 0 0 0 0 0 0 0
		b int64 // 1 1 1 1 1 1 1 1
		c int16 // 1 1 0 0 0 0 0 0
	}
	type S2 struct {
		a int8 // 1 0 1 1 0 0 0 0
		c int16
		b int64 // 1 1 1 1 1 1 1 1
	}
	fmt.Println(unsafe.Sizeof(S1{}), unsafe.Sizeof(S2{})) // 24 16

	type S3 struct {
		a int8 // 1 0 0 0 1 1 1 1
		b int32
		c int16 // 1 1 0 0
	}
	s3 := S3{}
	s1 := S1{}
	// 结构体大小是结构体对齐值的整数倍，结构体对齐值是结构体中最大字段的对齐值
	fmt.Println(unsafe.Alignof(s3), unsafe.Sizeof(s3), unsafe.Offsetof(s3.b), unsafe.Offsetof(s3.c)) // 4 12 4 8
	fmt.Println(unsafe.Alignof(s1), unsafe.Sizeof(s1))                                               // 8 24
}
