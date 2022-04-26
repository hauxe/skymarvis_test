package hau

import (
	"fmt"
)

// Assignment1
// direct the maxtrix index i,j from: left->right, top->bottom, right->left and bottom->top
func Assignment1(n int) {
	if n <= 0 {
		return
	}
	left, right, top, bottom := 0, n-1, 1, n-1
	direction := 1
	var i, j int
	defer fmt.Println()
	for {
		fmt.Print(i*n+j+1, " ")
		switch direction {
		case 1:
			if j == right {
				direction = 2
				right--
				i++
				if i > bottom {
					return
				}
			} else {
				j++
			}
		case -1:
			if j == left {
				direction = -2
				left++
				i--
				if i < top {
					return
				}
			} else {
				j--
			}
		case 2:
			if i == bottom {
				direction = -1
				bottom--
				j--
				if j < left {
					return
				}
			} else {
				i++
			}
		case -2:
			if i == top {
				direction = 1
				top++
				j++
				if j > right {
					return
				}
			} else {
				i--
			}
		}
	}
}
