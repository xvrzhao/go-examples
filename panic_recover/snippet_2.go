package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
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
