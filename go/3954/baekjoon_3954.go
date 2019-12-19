// https://www.acmicpc.net/problem/3954

package main

import "fmt"

type pair struct {
	start int
	end   int
}

func main() {
	var t int
	var result []string
	fmt.Scan(&t)
	result = make([]string, t)
	for i := 0; i < t; i++ {
		var sm, sc, si int
		var oper, input string
		var mem []byte
		var loopStack []int
		var loopPair []pair
		var idx = int(0)
		var inputIdx = int(0)
		var jmpPair pair
		jmpPair.start = -1
		jmpPair.end = -1
		j := 0
		fmt.Scan(&sm)
		mem = make([]byte, sm)
		fmt.Scan(&sc)
		loopPair = make([]pair, sc)
		fmt.Scan(&si)
		fmt.Scan(&oper)
		fmt.Scan(&input)
		for j, c := range oper {
			if c == '[' {
				loopStack = append(loopStack, j)
			}
			if c == ']' {
				loopStack = append(loopStack, j)
				top := len(loopStack) - 1
				var tmp pair
				tmp.start = loopStack[top-1]
				tmp.end = loopStack[top]
				loopPair = append(loopPair, tmp)
				loopStack = loopStack[:top-1]
			}
		}
		infiniteLoop := true
		for count := 0; count < 50000000; count++ {
			c := oper[j]
			if c == '-' {
				mem[idx]--
				j++
			} else if c == '+' {
				mem[idx]++
				j++
			} else if c == '<' {
				idx--
				if idx < 0 {
					idx = sm - 1
				}
				j++
			} else if c == '>' {
				idx++
				if idx >= sm {
					idx = 0
				}
				j++
			} else if c == '[' {
				if mem[idx] == 0 {
					for ii := range loopPair {
						if loopPair[ii].start == j {
							j = loopPair[ii].end
							break
						}
					}
				} else {
					j++
				}
			} else if c == ']' {
				if mem[idx] != 0 {
					for ii := range loopPair {
						if loopPair[ii].end == j {
							j = loopPair[ii].start
							break
						}
					}
				} else {
					j++
				}
			} else if c == '.' {
				j++
			} else if c == ',' {
				if inputIdx < si {
					mem[idx] = input[inputIdx]
					inputIdx++
				} else {
					mem[idx] = 255
				}
				j++
			}
			if j >= len(oper) {
				infiniteLoop = false
				break
			}
			if jmpPair.end < j {
				jmpPair.end = j
			}
		}
		if infiniteLoop {
			for _, p := range loopPair {
				if p.end == jmpPair.end {
					jmpPair.start = p.start
				}
			}
			result[i] = fmt.Sprintf("Loops %d %d\n", jmpPair.start, jmpPair.end)
		} else {
			result[i] = "Terminates\n"
		}
	}
	for _, r := range result {
		fmt.Printf(r)
	}
}
