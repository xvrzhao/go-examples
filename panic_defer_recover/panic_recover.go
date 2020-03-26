package panic_defer_recover

import (
	"errors"
	"fmt"
	"log"
)

func RunSnippet1() {
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

func RunSnippet2() {
	run()
}

func run() {
	s := make([]string, 3)
	if v, e := elementValue(s, 5); e != nil {
		log.Printf("Error: %s", e)
	} else {
		fmt.Printf("value is %v", v)
	}
}

// general function to get element value of a slice with no panic
func elementValue(slice []string, index int) (value string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%s", e))
		}
	}()
	value = slice[index]
	return
}

// ReturnedInterfaceValue demonstrates that the underlying type of
// the interface value returned by recover() is not error.
func ReturnedInterfaceValue() {
	defer func() {
		if v := recover(); v != nil {
			if err, ok := v.(error); ok == true {
				fmt.Printf("error: %v\n", err) // will not print
			} else {
				fmt.Printf("interface: %v\n", v) // will print
			}
		}
	}()
	panic("wrong!")
}
