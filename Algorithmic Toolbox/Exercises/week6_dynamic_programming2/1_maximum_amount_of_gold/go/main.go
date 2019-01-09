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

func getMaximumAmountOfGold(W, n int, arrayGold []int) [][]int {
	dpMax := make([][]int, n+1)

	for i := range dpMax {
		dpMax[i] = make([]int, W+1)
	}

	for i := 1; i < n+1; i++ {
		for w := 1; w < W+1; w++ {
			dpMax[i][w] = dpMax[i-1][w]

			if arrayGold[i-1] <= w {
				gold := dpMax[i-1][w - arrayGold[i-1]] + arrayGold[i-1]
				if dpMax[i][w] < gold {
					dpMax[i][w] = gold
				}
			} 

		}
	}

	return dpMax
}

func main() {

	defer writer.Flush()
	
	var n, W int

	scanf("%d %d", &W, &n)
	reader.ReadString('\n')
	
	arrayGold := make([]int, n)

	for index := 0; index < n; index++ {
		scanf("%d ", &arrayGold[index])
	}
	result := getMaximumAmountOfGold(W, n, arrayGold)
	printf("%v", result)
}