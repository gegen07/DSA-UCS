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

func min(arr []int) int {
	minor := 2000
	for _, x := range arr {
		if x < minor {
			minor = x
		}
	}

	return minor
}

// Time Complexity O(mn) 
// m is the money and n is the len of coins 
func getMoneyChange(money int, coins []int) 	int {
	var ( 
		memo []int
		arrNumCoins []int
	)

	memo = make([]int, money+1)
	for i := 1; i <= money; i++ {
		arrNumCoins = []int{}
		for _, coin := range coins {
			if i >= coin {
				arrNumCoins = append(arrNumCoins, (memo[i-coin]+1))
			}
		}
		memo[i] = min(arrNumCoins)
	}

	return memo[money]
}

func main() {
	defer writer.Flush()

	var n int

	scanf("%d\n", &n)

	cents := []int{1, 3, 4}

	result := getMoneyChange(n, cents)

	printf("%v", result)
}