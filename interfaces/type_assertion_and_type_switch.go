package interfaces

import "fmt"

type Human interface {
	Speak()
}

type Man struct {
	Name string
}

func (m *Man) Speak() {
	fmt.Printf("Hello, my name is %s", m.Name)
}

type Girl struct {
	Gender string
}

func (g *Girl) Speak() {}

func NewHuman(name string) Human {
	return &Man{Name: name}
}

func RunTypeAssertionExample() {
	human := NewHuman("Xavier")

	/*
	 * unchecked type assertion
	 */

	// girl := human.(*Girl) // panic as human is not *Girl
	// man := human.(*Man) // man, the value of type assertion is type *Man

	/*
	 * checked type assertion, which is the best practice
	 */

	girl, ok := human.(*Girl) // not panic although human is not *Girl, girl is type *Girl and value is `zero value` of *Girl, ok is false
	if ok {
		fmt.Println(girl.Gender)
	} else {
		fmt.Println("human is not *Girl")
	}

	// man is type *Man that unwrapped from human interface, ok is true
	if man, ok := human.(*Man); ok {
		fmt.Println(man.Name)
	} else {
		fmt.Println("human is not *Man")
	}

	// just to judge
	if _, ok = human.(*Girl); ok {
		// ...
	}
}

func RunTypeSwitchExample() {
	human := NewHuman("Xavier")

	switch t := human.(type) {
	case *Man:
		fmt.Println("*Man", t.Name)
	case *Girl:
		fmt.Println("*Girl", t.Gender)
	default:
		fmt.Println("human is not *Man or *Girl")
	}

	// fmt.Println(t) // can not use t outside the switch scope

	switch human.(type) {
	case *Man:
		// todo
	case *Girl:
		// todo
	default:
		// todo
	}
}
