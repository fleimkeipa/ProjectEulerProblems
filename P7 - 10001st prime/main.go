package main

import (
	"fmt"
	"time"
)

func Prime(a int) bool {
	for i := 2; i < a/2; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func FindPrimeWithForce(a int) int {
	temp := 2
	sayac := 0
	for {
		if sayac == a {
			break
		}
		temp++
		if Prime(temp) {
			sayac++
		}
	}
	return temp
}

func FindPrimeWithEratostenSieve(n int) int {
	prime := make(map[int]int, n)
	temp := 0
	for {
		for p := 2; p*p <= n; p++ {
			if prime[p] == 0 {
				temp++
				for i := p * p; i <= n; i += p {
					prime[i] = 1
				}
			}
		}

		for i := 0; i <= n; i++ {
			if prime[i] == 0 {
				temp++
			}
			if temp > 10068 {
				return i
			}
		}
	}
}

func main() {
	start := time.Now()
	// fmt.Println(FindPrimeWithForce(10001))
	fmt.Println(FindPrimeWithEratostenSieve(104744))
	fmt.Println(time.Since(start))
}
