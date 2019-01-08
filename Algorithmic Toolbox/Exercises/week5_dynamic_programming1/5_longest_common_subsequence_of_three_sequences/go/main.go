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

func max(x ... int) int {
	maxValue := x[0]

	for _, value := range x {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

/* LCS - Longest common subsequence Algorithm
  Using three sequences
*/

func LCS(x1, x2, x3 []int, n1, n2, n3 int) int {

	lcs := make([][][]int, n1)

	for i := range lcs {
		lcs[i] = make([][]int, n2)
		for j := range lcs[i] {
			lcs[i][j] = make([]int, n3)
		}
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			for k := 0; k < n3; k++ {
				if i == 0 || j == 0 || k == 0 {
					lcs[i][j][k] = 0
				} else if x1[i-1] == x2[j-1] &&
						  x1[i-1] == x3[k-1] {
					lcs[i][j][k] = lcs[i-1][j-1][k-1] + 1
				} else {
					lcs[i][j][k] = max(lcs[i-1][j][k], lcs[i][j-1][k], lcs[i][j][k-1])
				}
			}
		}
	}

	return lcs[n1-1][n2-1][n3-1]
}

func main() {

	defer writer.Flush()
	
	var (
		x1, x2, x3 []int
		n1, n2, n3 int
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

	scanf("\n%d\n", &n3)

	x3 = make([]int, n3)
	for i := 0; i < n3; i++ {
		scanf("%d ", &x3[i])
	}

	lcs := LCS(x1, x2, x3, n1+1, n2+1, n3+1)
	printf("%v\n", lcs)
}