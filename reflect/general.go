package reflect

import (
	"fmt"
	"reflect"
)

func RunSetExample() {
	f := 3.14

	fv := reflect.ValueOf(&f)
	if fv.Kind() == reflect.Ptr {
		fv = fv.Elem()
	}
	if fv.CanSet() && fv.Kind() == reflect.Float64 {
		fv.SetFloat(3.1415)
	}

	fmt.Println(f) // 3.1415
}

type person struct {
	Name   string
	Age    uint8
	gender uint8
}

func RunSetStructExample() {
	v := reflect.ValueOf(person{})
	fmt.Println(v.CanSet()) // false

	v = reflect.ValueOf(&person{})
	fmt.Println(v.CanSet()) // false

	v = reflect.ValueOf(&person{}).Elem()
	fmt.Println(v.CanSet()) // true

	p := person{}
	v = reflect.ValueOf(&p).Elem()
	nv := v.FieldByName("Name")
	fmt.Println(nv.Kind(), nv.CanSet()) // string true
	nv.SetString("Xavier")
	fmt.Println(p.Name) // Xavier

	gv := v.FieldByName("gender")
	fmt.Println(gv.Kind(), gv.CanSet()) // uint8 false(gender is unexported)

	zv := v.FieldByName("notExistField") // sv is zero Value
	// fmt.Println(sv.IsZero()) // panic: reflect: call of reflect.Value.IsZero on zero Value
	fmt.Println(zv.IsValid(), zv.CanSet()) // false false
}
