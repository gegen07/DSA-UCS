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

func getMaximumNumberPrizes(number, controller int) []int {
	var array []int

	for number > 2*controller {
		array = append(array, controller)
		number -= controller
		controller++	
	}

	array = append(array, number)

	return array
}

func main() {

	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	array := getMaximumNumberPrizes(n, 1)
	
	printf("%d\n", len(array))

	for _, position := range array {
		printf("%d ", position)
	}
}