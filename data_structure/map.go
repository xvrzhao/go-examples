package data_structure

import "fmt"

func RunMapPanic() {
	// m := make(map[string]int) // declared and initialized
	var m map[string]int  // declared but not initialized
	fmt.Println(m == nil) // true
	m["xvrzhao"] = 23     // panic: assignment to entry in nil map
}

func RunMapRefer() {
	m1 := make(map[string]string, 2)
	m1["apple"] = "red"
	m1["banana"] = "yellow"

	m2 := m1 // map is reference type, m1 and m2 point to the same block of memory
	m2["apple"] = "green"

	fmt.Println(m1) // map[apple:green banana:yellow]
	fmt.Println(m2) // map[apple:green banana:yellow]
}

func RunMapExceededCap() {
	m1 := make(map[string]string, 2)
	m1["apple"] = "red"
	m1["banana"] = "yellow"

	m2 := m1 // m2 points to m1, or m2 and m1 point to the same block of memory

	m1["orange"] = "orange" // m2 also changes, even though m1 has exceeded capacity limit
	fmt.Println(m1)         // map[apple:red banana:yellow orange:orange]
	fmt.Println(m2)         // map[apple:red banana:yellow orange:orange]
}

type Age struct {
	value int
}

// value of the key in map is not addressable, so when value stores the struct type or array type,
// you can not change element of that value directly. So, usually we set value's type in map to struct pointer,
// array pointer, slice, or map.
func RunMapNotAddressable() {
	m := map[string]Age{
		"xavier": Age{value: 23},
		"john":   Age{value: 13},
	}

	//m["xavier"].value = 10 // compile error: cannot assign to struct field m["xavier"].value in map

	// change value of the key in map
	xavier := m["xavier"]
	xavier.value = 10
	m["xavier"] = xavier

	// or the following function
}

func RunMapAddressable() {
	// value of the key in map stores pointer of the struct
	m := map[string]*Age{
		"xavier": &Age{value: 23},
	}

	m["xavier"].value = 10
	fmt.Println(m["xavier"])
}
