package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPairwiseProductNaive(inputNumbers []int) int {
	var product int
	for _, i := range inputNumbers {
		for _, j := range inputNumbers {
			if i != j {
				product = Max(product, i*j)
			}
		}
	}

	return product
}

func maxPairwiseProductFast(inputNumbers []int) int {
	index1 := 0

	for index := 1; index < len(inputNumbers); index++ {
		if inputNumbers[index] > inputNumbers[index1] {
			index1 = index
		}
	}

	var index2 int
	
	if index1 == 0 {
		index2 = 1
	} else {
		index2 = 0
	}
	
	for index := 1; index < len(inputNumbers); index++ {
		if inputNumbers[index] > inputNumbers[index2] && index != index1 {
			index2 = index
		}
	}

	return inputNumbers[index1] * inputNumbers[index2]
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	array := make([]int, n)

	for index := 0; index < n; index++ {
		fmt.Scanf("%d", &array[index])
	}

	var result = maxPairwiseProductFast(array)
	fmt.Printf("%v", result)
}
