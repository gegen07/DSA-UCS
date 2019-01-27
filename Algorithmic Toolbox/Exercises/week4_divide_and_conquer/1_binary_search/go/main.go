package main 

import (
	"fmt"
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func binarySearch(array []int, low, high, key int) int {
	if high <= low {
		return -1
	}

	mid := low + int(((high-low)/2))
	if key == array[mid] {
		return mid
	} else if key > array[mid]{
		return binarySearch(array, mid+1, high, key)
	} else {
		return binarySearch(array, low, mid-1, key)
	}
}

func main() {

	defer writer.Flush()

	var n int 

	scanf("%d ", &n)

	array := make([]int, n)

	for index := 0; index < n; index++ {
		scanf("%d ", &array[index])
	}

	scanf("\n")
	scanf("%d ", &n)

	searching := make([]int, n)

	for index := 0; index < n; index++ {
		scanf("%d ", &searching[index])
	}

	for index := 0; index < n; index++ {
		printf("%d ", binarySearch(array, 0, len(array), searching[index]))
	}
}