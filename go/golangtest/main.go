package main

import "fmt"

func main() {
	var i = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var x = append([]int{}, i...)
	x[5] = 500000
	for idx := range i {
		fmt.Printf("%d ", i[idx])
	}
	fmt.Printf("\n")
	for idx := range x {
		fmt.Printf("%d ", x[idx])
	}
	fmt.Printf("\n")
}
