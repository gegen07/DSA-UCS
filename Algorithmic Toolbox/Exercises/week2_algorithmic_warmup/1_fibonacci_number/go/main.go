package main

import "fmt"

func fibonacciNaive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciNaive(n-1)+fibonacciNaive(n-2)
}

func fibonacciFast(n int) int {
	x := make([]int, n)

	x[0] = 0
	x[1] = 1

	for index := 2; index < n; index++ {
		x[index] = x[index-1] + x[index-2]
	}

	return x[n]
}

func main () {
	var n int
	fmt.Scanf("%d", &n);

	fmt.Printf("%v", fibonacciNaive(n))
}

