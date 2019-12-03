// https://www.acmicpc.net/problem/17406

package main

import "fmt"

func sum_mat(a [][]int) int {
	var sum int = 0
	l := len(a)
	for i := 0; i < l; i++ {
		ll := len(a[i])
		var tmp int = 0
		for j := 0; j < ll; j++ {
			tmp = tmp + a[i][j]
		}
		if sum > tmp {
			sum = tmp
		}
	}
	return sum
}

func main() {
	var n, m, k int
	fmt.Scanf("%d %d %d", &n, &m, &k)
	for i := 0; i < n; i++ {

	}
}
