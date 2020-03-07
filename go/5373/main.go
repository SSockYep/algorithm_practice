// https://www.acmicpc.net/problem/5373

package main

import "fmt"

type cube struct {
	u, d, f, b, l, r [3][3]byte
}

func rotate(c *cube, face byte, clockwise bool) {
	if clockwise {
		switch face {
		case 'u':
		case 'd':
		case 'f':
		case 'b':
		case 'l':
		case 'r':
		}
	} else {
		switch face {
		case 'u':
			tmp := c.f[0]
			c.f[0] = c.l[0]
			c.l[0] = c.b[0]
			c.b[0] = c.r[0]
			c.r[0] = tmp
		case 'd':
		case 'f':
		case 'b':
		case 'l':
		case 'r':
		}

	}
}

func main() {
	var a cube
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a.u[i][j] = 'w'
			a.d[i][j] = 'y'
			a.f[i][j] = 'f'
			a.b[i][j] = 'o'
			a.l[i][j] = 'g'
			a.r[i][j] = 'b'
		}
	}

	for _, v := range a.u {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		fmt.Printf("\n")
	}
}
