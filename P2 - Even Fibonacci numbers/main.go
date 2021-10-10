package main

import "fmt"

func fiboc(come int) int {
	first, second := 1, 1
	sumEven := 0
	for {
		if second < come {
			first, second = second, first+second
			if second%2 == 0 {
				sumEven += second
			}
		} else {
			break
		}
	}
	return sumEven
}

func main() {
	fmt.Println(fiboc(4000000))
}
