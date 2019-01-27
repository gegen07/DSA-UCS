package main

import "fmt"

func findLastDigitFibonacci (n int) int {
	x := make([]int, n)

	x[0] = 0
	x[1] = 1

	for index := 2; index < n; index++ {
		x[index] = (x[index-1] + x[index+2])%10
	}
	return x[n]
}

func main() {
	var n int

	fmt.Scanf("%v", &n)

	fmt.Printf("%v", findLastDigitFibonacci(n))
}