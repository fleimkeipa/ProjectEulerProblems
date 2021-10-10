package main

import (
	"fmt"
	"math/big"
	"time"
)

var sayac int

func Prime(come, max int) int {
	for i := come + 1; i < max/2; i++ { //come = 5, i=5,6,7,8
		sayac++
		if i%2 == 0 {
			continue
		}
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			return i
		}
	}
	return max
}

func LargestPrime(come int) int {
	prime1 := 2
	max := 0
	for {
		if come > prime1 {
			if come%prime1 == 0 {
				max = prime1
				come /= prime1
				if big.NewInt(int64(come)).ProbablyPrime(0) {
					return come
				}
			}
			prime1 = Prime(prime1, come)
		} else {
			break
		}
	}
	return max
}

func largestPrime2(come int) int {
	prime1 := 2
	max := 0
	for come > 1 {
		if come%prime1 == 0 {
			max = prime1
			come /= prime1
		} else {
			prime1 = Prime(prime1, come)
		}
	}
	return max
}

func main() {
	start := time.Now()
	fmt.Println(LargestPrime(600851475143)) // duration = 3.877595ms 	sayac = 1469
	//fmt.Println(largestPrime2(600851475143)) 	// duration = 10.012578ms	sayac = 3425
	fmt.Println(time.Since(start))
	fmt.Println("sayac", sayac)
}
