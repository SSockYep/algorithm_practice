package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sn := make([]int, len(s))
	copy(sn, s)
	for _, c := range sn {
		fmt.Printf("%d ", c)
	}
	fmt.Printf("\n")
}
