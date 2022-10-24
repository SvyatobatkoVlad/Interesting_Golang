package main

import "fmt"

func main() {
	s := 1
	ints := []int{1, 2}
	fmt.Println(Contains(s, ints))

	strs := []string{"1", "2"}
	PrintSlice(ints)
	PrintSlice(strs)
}

func Contains[T comparable](v T, vv []T) bool {
	for _, a := range vv {
		if a == v {
			return true
		}
	}
	return false
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
