// https://www.acmicpc.net/problem/16639

package main

import (
	"fmt"
)

type sign struct {
	s    byte
	prio int
}

const maxInt = int(^uint(0) >> 1)
const minInt = ^maxInt

var ans int

func calc(l int, s sign, r int) int {
	if s.s == '*' {
		return l * r
	}
	if s.s == '+' {
		return l + r
	}
	if s.s == '-' {
		return l - r
	}
	return 0
}

func calcExp(numbers []int, signs []sign) int {
	if len(numbers)-len(signs) != 1 {
		return maxInt
	}
	newNumbers := append([]int{}, numbers...)
	newSigns := append([]sign{}, signs...)

	for len(newNumbers) > 1 {
		maxPrio := minInt
		maxPrioIdx := minInt
		for i, v := range newSigns {

			if v.prio > maxPrio {
				maxPrio = v.prio
				maxPrioIdx = i
			}
		}
		c := calc(newNumbers[maxPrioIdx], newSigns[maxPrioIdx], newNumbers[maxPrioIdx+1])
		newNumbers = append(append(newNumbers[:maxPrioIdx], c), newNumbers[maxPrioIdx+2:]...)
		newSigns = append(newSigns[:maxPrioIdx], newSigns[maxPrioIdx+1:]...)
	}
	return newNumbers[0]
}
func deleteAttrib(arr []int, a int) []int {
	newArr := append([]int{}, arr...)

	return newArr
}
func search(numbers []int, signs []sign, n int, p int, r []int) {
	newSigns := append([]sign{}, signs...)

	var newRest = append([]int{}, r...)
	if n >= 0 {
		newSigns[n].prio = p
		newRest[p] = -1
	}
	if n == len(signs)-1 {
		tmp := calcExp(numbers, newSigns)
		if tmp > ans {
			ans = tmp
		}
		return
	}

	for _, v := range newRest {
		if v != -1 {
			search(numbers, newSigns, n+1, v, newRest)
		}
	}

}

func main() {
	var n int
	var exp string
	var numbers []int
	var signs []sign
	ans = minInt
	fmt.Scan(&n)
	fmt.Scan(&exp)
	numbers = make([]int, (n+1)/2)
	signs = make([]sign, (n-1)/2)
	perm := make([]int, len(signs))

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			numbers[i/2] = int(exp[i] - '0')
		} else {
			signs[(i-1)/2] = sign{exp[i], -1}
		}
	}
	for i := range perm {
		perm[i] = i
	}

	search(numbers, signs, -1, -1, perm)
	fmt.Printf("%d\n", ans)
}
