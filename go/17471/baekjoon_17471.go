// https://www.acmicpc.net/problem/17471

package main

import (
	"fmt"
)

const max int = int(^(uint(0)) >> 1)

func checkConnect(a []bool, connect [][]bool) bool {
	as := []int{}
	bs := []int{}
	for i := range a {
		if a[i] {
			as = append(as, i)
		} else {
			bs = append(bs, i)
		}
	}
	return bfs(as, connect) && bfs(bs, connect)
}

func calcPopulation(a []bool, pop []int) int {
	var aSum, bSum int
	aSum = 0
	bSum = 0
	for i := range a {
		if a[i] {
			aSum += pop[i]
		} else {
			bSum += pop[i]
		}
	}
	ret := aSum - bSum
	if ret >= 0 {
		return ret
	}
	return -ret
}

func search(wardPop []int, connect [][]bool) int {
	var a []bool
	ret := max
	a = make([]bool, len(wardPop))
	for bitCounter := uint(0); bitCounter < 1<<uint(len(wardPop)); bitCounter++ {
		for i := 0; i < len(wardPop); i++ {
			if ((1<<uint(i))&bitCounter)>>uint(i) == 1 {
				a[i] = true
			} else {
				a[i] = false
			}
		}

		if checkConnect(a, connect) {
			calc := calcPopulation(a, wardPop)
			if ret > calc {
				ret = calc
			}
		}
	}
	if ret == max {
		return -1
	}
	return ret
}

func bfs(a []int, connect [][]bool) bool {
	if len(a) == 0 {
		return false
	}
	searched := make([]bool, len(connect))
	searched[a[0]] = true
	queue := []int{a[0]}
	for len(queue) > 0 {
		var u = queue[0]
		queue = append(queue[:0], queue[1:]...)
		for j := range connect[u] {
			if connect[u][j] && searched[j] == false {
				var flag = false
				for _, value := range a {
					if value == j {
						flag = true
					}
				}
				if flag {
					searched[j] = true
					queue = append(queue, j)
				}
			}
		}
	}
	for _, v := range a {
		if !searched[v] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var connect [][]bool
	var wardPop []int
	fmt.Scan(&n)
	connect = make([][]bool, n)
	wardPop = make([]int, n)
	for i := range connect {
		connect[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&wardPop[i])
	}
	for i := 0; i < n; i++ {
		var cn int
		fmt.Scan(&cn)
		for j := 0; j < cn; j++ {
			var tmp int
			fmt.Scan(&tmp)
			connect[i][tmp-1] = true
		}
	}
	fmt.Printf("%d\n", search(wardPop, connect))
}
