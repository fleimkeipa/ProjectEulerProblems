package main

import "fmt"

func Multiples(a, b, c int) int {
	sum := 0
	for i := 0; i < c; i++ {
		if i%a == 0 || i%b == 0 {
			sum += i
			fmt.Println(i)
		}
	}
	return sum
}

func main() {
	fmt.Println(Multiples(3, 5, 1000))
}
