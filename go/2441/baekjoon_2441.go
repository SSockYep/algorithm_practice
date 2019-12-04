// https://www.acmicpc.net/problem/2441

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := n; i > 0; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Printf(" ")
		}
		for j := i; j > 0; j-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
