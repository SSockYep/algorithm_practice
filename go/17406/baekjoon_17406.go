// https://www.acmicpc.net/problem/17406

package main

import "fmt"

var did []bool
var exps []exp
var ans int

type exp struct {
	r int
	c int
	s int
}

func sumMat(a [][]int) int {
	var sum int = int(^uint(0) >> 1)
	for i := 0; i < len(a); i++ {
		var tmp int = 0
		for j := 0; j < len(a[i]); j++ {
			tmp = tmp + a[i][j]
		}
		if sum > tmp {
			sum = tmp
		}
	}
	return sum
}

func calc(mat [][]int, e exp) [][]int {
	var _mat = make([][]int, len(mat), len(mat))
	for i := 0; i < len(_mat); i++ {
		_mat[i] = make([]int, len(mat[i]), len(mat[i]))
		copy(_mat[i], mat[i])
	}
	var s = e.s
	var r = e.r
	var c = e.c
	for i := s; i > 0; i-- {
		_mat[r-i][c-i] = mat[r-i+1][c-i]
		_mat[r-i][c+i] = mat[r-i][c+i-1]
		_mat[r+i][c-i] = mat[r+i][c-i+1]
		_mat[r+i][c+i] = mat[r+i-1][c+i]
		for j := c - i + 1; j < c+i; j++ {
			_mat[r-i][j] = mat[r-i][j-1]
			_mat[r+i][j] = mat[r+i][j+1]
		}
		for j := r - i + 1; j < r+i; j++ {
			_mat[j][c-i] = mat[j+1][c-i]
			_mat[j][c+i] = mat[j-1][c+i]
		}
	}
	return _mat
}

func dfs(idx int, mat [][]int) {
	var m = mat
	if idx != -1 {
		m = calc(mat, exps[idx])
	}
	var flag = true
	for i := 0; i < len(did); i++ {
		if did[i] == false {
			flag = false
			break
		}
	}
	if flag {
		var sum = sumMat(m)
		if sum < ans {
			ans = sum
		}
	} else {
		for i := 0; i < len(did); i++ {
			if !did[i] {
				did[i] = true
				dfs(i, m)
				did[i] = false
			}
		}
	}
}

func main() {
	var n, m, k int
	var mat [][]int
	var arr []int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	mat = make([][]int, n, n)
	exps = make([]exp, k, k)
	did = make([]bool, k, k)
	ans = int(^uint(0) >> 1)

	for i := 0; i < n; i++ {
		arr = make([]int, m, m)
		mat[i] = arr
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var tmp int
			fmt.Scan(&tmp)
			mat[i][j] = tmp
		}
	}

	for i := 0; i < k; i++ {
		var tr, tc, ts int
		fmt.Scan(&tr)
		fmt.Scan(&tc)
		fmt.Scan(&ts)
		exps[i].r = tr - 1
		exps[i].c = tc - 1
		exps[i].s = ts
	}
	dfs(-1, mat)

	fmt.Printf("%d", ans)

}
