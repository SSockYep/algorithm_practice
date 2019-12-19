// https://www.acmicpc.net/problem/16639

package main

import (
	"fmt"
)

type pair struct {
	s, e int
	b    bool
}

var resmap map[pair]int
var numbers []int
var signs []byte

func f(l int, s byte, r int) int {
	switch s {
	case '+':
		return l + r
	case '-':
		return l - r
	case '*':
		return l * r
	default:
		return l + r
	}
}

func calc(s int, e int, b bool) int {
	ret := int(^(uint(0)) >> 1)
	if b {
		ret = -ret
	}
	var p pair
	var tmp int
	p.s = s
	p.e = e
	p.b = b
	v, exist := resmap[p]
	if exist {
		return v
	}
	if s > e {
		return 0
	}
	if s == e {
		resmap[p] = numbers[s]
		return resmap[p]
	}
	if e-s == 1 {
		tmp := f(numbers[s], signs[s], numbers[e])
		resmap[p] = tmp
		return resmap[p]
	}
	for i := 0; i < e-s; i++ {
		for j := s; j+i <= e; j++ {

		}
	}
	resmap[p] = tmp
	return ret
}

func main() {
	var n int
	var equation string

	resmap = make(map[pair]int)

	fmt.Scan(&n)
	fmt.Scan(&equation)
	numbers = make([]int, (n+1)/2)
	signs = make([]byte, (n-1)/2)

	for i, c := range equation {
		if i%2 == 0 {
			numbers[i/2] = int(c - '0')
		} else {
			signs[(i-1)/2] = byte(c)
		}
	}

	fmt.Printf("%d", calc(0, len(numbers)-1, true))

}
