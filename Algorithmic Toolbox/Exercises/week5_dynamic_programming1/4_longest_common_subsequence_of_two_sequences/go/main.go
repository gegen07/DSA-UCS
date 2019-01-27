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

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	
	return x2
}

//LCS - Longest common subsequence Algorithm
func LCS(x1, x2 []int, n1, n2 int) int {

	lcs := make([][]int, n1)

	for i := range lcs {
		lcs[i] = make([]int, n2)
	} 

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if x1[i-1] == x2[j-1]{
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	return lcs[n1-1][n2-1]
}

func main() {

	defer writer.Flush()
	
	var (
		x1, x2 []int
		n1, n2 int
	)
	scanf("%d\n", &n1)

	x1 = make([]int, n1)
	for i := 0; i < n1; i++ {
		scanf("%d ", &x1[i])
	}
	
	scanf("\n%d\n", &n2)
	
	x2 = make([]int, n2)
	for i := 0; i < n2; i++ {
		scanf("%d ", &x2[i])
	}

	lcs := LCS(x1, x2, n1+1, n2+1)
	printf("%v\n", lcs)
}