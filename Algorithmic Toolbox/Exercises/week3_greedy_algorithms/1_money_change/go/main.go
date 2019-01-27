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

func getMoneyChange(n int32) int32 {
	cents := []int32{10, 5, 1}
	var result int32
	
	for index := 0; index < len(cents); index++ {
		cent := cents[index]
		times := (n-n%cent)/cent
		n -= times * cent
		result += times
	}

	return result
}

func main() {
	var n int32

	fmt.Scanf("%v\n", &n)

	result := getMoneyChange(n)

	fmt.Printf("%v", result)

	defer writer.Flush()


}