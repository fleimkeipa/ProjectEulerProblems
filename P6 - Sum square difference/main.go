package main

import "fmt"

func SumSquareDifference(a int) int {
	sum1 := (a * a) * (a + 2) * (a + 2) / 4
	sum2 := a * (a + 1) * (2*a + 1) / 6
	return sum1 - sum2
}

func main() {
	fmt.Println(SumSquareDifference(100))
}
