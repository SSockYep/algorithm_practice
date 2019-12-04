// https://www.acmicpc.net/problem/2443

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
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
