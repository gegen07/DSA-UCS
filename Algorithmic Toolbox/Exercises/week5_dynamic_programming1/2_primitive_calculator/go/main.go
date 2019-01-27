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

func min(x1, x2 int) int {
	if x1 <= x2 {
		return x1
	}
	return x2
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func getPrimitiveCalculatorSequenceAux(n int) []int {
	arr := make([]int, n+1)

	for index := 1; index <= n; index++ {
		arr[index] = arr[index-1] + 1

		if index % 2 == 0 {
			arr[index] = min(arr[index/2]+1, arr[index])
		}
		
		if index % 3 == 0 {
			arr[index] = min(arr[index/3]+1, arr[index])
		}
	}
	return arr
}

func getPrimitiveCalculatorSequence(n int) []int {
	arr := getPrimitiveCalculatorSequenceAux(n)
	
	var sequence []int
	
	for n > 1 {
		sequence = append(sequence, n)

		if arr[n-1] == arr[n]-1 {
			n--
		} else if n % 2 == 0 && arr[n/2] == arr[n]-1 {
			n /= 2
		} else if n % 3 == 0 && arr[n/3] == arr[n]-1 {
			n /= 3
		}
	}

	sequence = append(sequence, 1)
	reverse(sequence)
	return sequence
}

func main() {

	defer writer.Flush()

	var n int

	scanf("%d\n", &n)

	sequence := getPrimitiveCalculatorSequence(n)
	printf("%v\n%v\n", len(sequence)-1, sequence)
}