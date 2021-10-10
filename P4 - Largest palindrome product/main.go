package main

import (
	"fmt"
	"time"
)

func Palindrome(a int) bool {
	temp := a
	reverse := 0
	for temp > 0 {
		reverse *= 10
		reverse += (temp % 10)
		temp /= 10
	}
	if a == reverse {
		return true
	}
	return false
}
func LargestPalindrome() int {
	max := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if Palindrome(i * j) {
				if max < i*j {
					max = i * j
				}
				break
			}
		}
	}
	return max
}

func main() {
	start := time.Now()
	fmt.Println(LargestPalindrome())
	fmt.Println(time.Since(start))
}
