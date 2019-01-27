package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func getMaximumAdvertisementRevenue(arr1, arr2 []int, n int) int {
	if n > 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(arr1)))
		sort.Sort(sort.Reverse(sort.IntSlice(arr2)))
	}

	var sum int
	for index := 0; index < n; index++ {
		sum += arr1[index] * arr2[index]
	}

	return sum
}

func main() {

	defer writer.Flush()

	var n int
	scanf("%v\n", &n)
	
	arr1 := make([]int, n)
	arr2 := make([]int, n)
	for index := 0; index < n; index++ {
		scanf("%d ", &arr1[index])
	}
	scanf("\n")	

	for index := 0; index < n; index++ {
		scanf("%d ", &arr2[index])
	}

	printf("%v", getMaximumAdvertisementRevenue(arr1, arr2, n))
}