package main

import (
	"fmt"
	"time"
)

func PythagoranTriplet(n int) {
	for i := 2; i < (n-3)/3; i++ {
		for j := i + 1; j < (n-i-1)/2; j++ {
			k := 1000 - i - j
			if i*i+j*j == k*k && i+j+k == n {
				fmt.Println("i", i, "j", j, "k", k, "ijk", i*j*k)
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
