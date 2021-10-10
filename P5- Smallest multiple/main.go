package main

import (
	"fmt"
)

func SmallestMultiple(a, b int) int {
	c, temp := 1, 1
	for {
		for i := a; i <= b; i++ {
			if temp == b-1 {
				return c
			}
			if c%i == 0 {
				temp = i
			} else {
				break
			}
		}
		temp = 1
		c++
	}
}

func main() {
	fmt.Println(SmallestMultiple(1, 20))
}
