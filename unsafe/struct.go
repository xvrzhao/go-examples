package unsafe

import (
	"reflect"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

// Person2Bytes 将 Person 结构体转为 bytes。
// 这里或许不能称为转，应该为获取结构体的内存引用。
// 需要注意 GC 问题，如果结构体被回收了，根据 bytes 的 Data(起点) 和 Len(长度) 获取到的内存可能已经不再是之前的结构体了。
//
// TODO: Translate to English.
func Person2Bytes(p *Person) []byte {
	size := unsafe.Sizeof(*p)

	s := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  int(size),
		Cap:  int(size),
	}

	return *(*[]byte)(unsafe.Pointer(&s))
}

// Bytes2Person 将 bytes 再转回结构体。
// 结合 Person2Bytes 函数进行来回互转见单元测试：TestPerson2Bytes。项目根目录下执行：
//   go test -v -run=Person2Bytes ./unsafe
//
// TODO: Translate to English.
func Bytes2Person(b []byte) *Person {
	personAddress := (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	return (*Person)(unsafe.Pointer(personAddress))
}
