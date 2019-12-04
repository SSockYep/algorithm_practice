// https://www.acmicpc.net/problem/2444

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i < n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for j := 0; j < i*2-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i := n; i > 0; i-- {
		for j := n - i; j > 0; j-- {
			fmt.Printf(" ")
		}
		for j := 0; j < i*2-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

}
