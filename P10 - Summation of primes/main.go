package main

import (
	"fmt"
	"math"
	"time"
)

func SummationPrimeWithEratostenSieve(n int) int { //send 2000000
	prime := make(map[int]int, n)
	crossLimit := math.Sqrt(float64(n))
	for i := 4; i < n; i += 2 {
		prime[i] = 1
	}
	for p := 3; p <= int(crossLimit); p += 2 {
		if prime[p] == 0 {
			for i := p * p; i <= n; i += 2 * p {
				prime[i] = 1
			}
		}
	}
	sum := 2
	for i := 3; i <= n; i += 2 {
		if prime[i] == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	start := time.Now()
	fmt.Println(SummationPrimeWithEratostenSieve(2000000))
	fmt.Println(time.Since(start))
}
