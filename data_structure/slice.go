package data_structure

import "fmt"

func appendSlice1(s []int) {
	length := cap(s) - len(s) + 1
	suf := make([]int, length)
	s = append(s, suf...)
}

func appendSlice2(s *[]int) {
	length := cap(*s) - len(*s) + 1
	suf := make([]int, length)
	*s = append(*s, suf...)
}

func RunAppendSliceWithFunc() {
	s1 := []int{1, 2, 3}
	appendSlice1(s1)
	fmt.Println("s1:", s1)

	s2 := []int{1, 2, 3}
	appendSlice2(&s2)
	fmt.Println("s2:", s2)
}
