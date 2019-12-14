package interfaces

import "fmt"

type Human interface {
	Speak()
}

type Animal interface {
	Walk()
}

type Man struct {
	Name string
}

func (m *Man) Speak() {
	fmt.Printf("Hello, my name is %s", m.Name)
}

func (m Man) Walk() {
	fmt.Println("Walking...")
}

type Girl struct {
	Gender string
}

func (g *Girl) Speak() {}

func newHuman(name string) Human {
	return &Man{Name: name}
}
