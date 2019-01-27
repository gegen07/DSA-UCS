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

func min(x1, x2, x3 int) int {
	if x1 < x3 && x1 < x2 {
		return x1
	} else if x2 < x1 && x2 < x3 {
		return x2
	}

	return x3
}

// Wagner-Fischer Algorithm
func getMinimumEditDistance(s1, s2 []rune) int {
	n := len(s1) + 1
	m := len(s2) + 1

	dist := make([][]int, n)

	for i := range dist {
		dist[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		dist[i][0] = i
	}

	for j := 0; j < m; j++ {
		dist[0][j] = j
	}

	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			if s1[i-1] == s2[j-1] {
				dist[i][j] = dist[i-1][j-1]
			} else {
				dist[i][j] = min(1 + dist[i-1][j], 
								 1 + dist[i][j-1], 
								 1 + dist[i-1][j-1])
			}
		}
	}

	return dist[len(s1)][len(s2)]
}

func main() {

	defer writer.Flush()

	var s1, s2 string
	scanf("%s\n%s\n", &s1, &s2)

	minDistance := getMinimumEditDistance([]rune(s1), []rune(s2))
	printf("%v", minDistance)
}