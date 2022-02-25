package main

import (
	"fmt"
	"time"
)

func PythagoranTriplet(n int) {
	for i := 2; i < (n-3)/3; i++ {
		for j := i + 1; j < (n-i-1)/2; j++ {
			k := n - i - j
			if i*i+j*j == k*k && i+j+k == n {
				fmt.Printf("i: %d j: %d k: %d i+j+k: %d \nijk: %d\n", i, j, k, i+j+k, i*j*k)
				return
			}
		}
	}
}

func main() {
	start := time.Now()
	PythagoranTriplet(1000)
	fmt.Println(time.Since(start))
}
