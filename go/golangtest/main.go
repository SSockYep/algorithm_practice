package main

import "fmt"

type s struct {
	a, b int
}

func (a s) (v s) s {
	r := a
	r.a += v.a
	r.b += v.b
	return r
}

func main() {
	var l, r s
	l.a = 1
	l.b = 2
	r.a = 10
	r.b = 20
	var ans = l + r
	fmt.Printf("%d %d\n", ans.a, ans.b)
}
