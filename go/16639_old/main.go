// https://www.acmicpc.net/problem/16639

package main

import (
	"fmt"
)

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

func calc(numbers []int, signs []byte) int {
	ret := -int(^uint(0) >> 1)

	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	if len(numbers) == 2 {
		return f(numbers[0], signs[0], numbers[1])
	}
	// for문의 조건을 바꿔서 중복돼서 계산하는 부분을 없애야함 --> 어떻게?? --> dfs를 쓰면 될거같은데
	// 현재 상황에서 1+2+3+4가 있는 경우 (1+2)+(3+4)를 (1+2)먼저 1회, (3+4)먼저 1회 총 2회 계산함
	for i := 0; i < len(numbers)-1; i++ {
		var pCalc, tmp int
		for jj := 1; i+jj < len(numbers); jj++ {
			j := i + jj
			newNumbers := make([]int, len(numbers))
			copy(newNumbers, numbers)
			newSigns := make([]byte, len(signs))
			copy(newSigns, signs)
			if i == 0 && j == len(numbers)-1 {
				continue
			}
			pNumbers := newNumbers[i : j+1]
			pSigns := newSigns[i:j]
			pCalc = calc(pNumbers, pSigns)
			tmp = calc(append(append(newNumbers[:i], pCalc), newNumbers[j+1:]...),
				append(newSigns[:i], newSigns[j:]...))

			if tmp > ret {
				ret = tmp
			}
		}
	}
	return ret
}

func main() {
	var n int
	var equation string
	var numbers []int
	var signs []byte

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

	fmt.Printf("%d", calc(numbers, signs))

}
