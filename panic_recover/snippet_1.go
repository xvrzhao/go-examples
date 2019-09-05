package main

import "fmt"

func main() {
	a()
	fmt.Println("The program is finished.")
}

func a() {
	defer func() {
		fmt.Println("This message will be displayed.")
		if err := recover(); err != nil {
			fmt.Printf("Rceived panic: `%v`, but program continues.\n", err)
		}
	}()
	b()
	fmt.Println("This message will not be displayed.")
}

func b() {
	panic("panic occured from func b.")
	fmt.Println("This message will not be displayed.")
}
