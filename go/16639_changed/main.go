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

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func search(numbers []int, signs []sign) {
	if len(signs) == 0 {
		ans = numbers[0]
		return
	}
	arr := make([]int, len(signs))
	for i := range arr {
		arr[i] = i
	}

	for p := make([]int, len(arr)); p[0] < len(p); nextPerm(p) {
		permutation := getPerm(arr, p)
		for i := range permutation {
			signs[i].prio = permutation[i]
		}
		tmp := calcExp(numbers, signs)
		if tmp > ans {
			ans = tmp
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

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			numbers[i/2] = int(exp[i] - '0')
		} else {
			signs[(i-1)/2] = sign{exp[i], 0}
		}
	}

	search(numbers, signs)
	fmt.Printf("%d\n", ans)
}
