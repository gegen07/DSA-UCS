package main

import "fmt"

func AplusB(a,b int64) int64 {
	return a+b
}

func main () {
	var a, b int64
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%v", AplusB(a, b))
}