package main

import "fmt"

func getPisanoPeriodNaive(n int) []int  {
	var period []int

	period = append(period, 0)
	period = append(period, 1)

	var remainder int

	for index := 0; index < n*n; index++ {
		remainder = (period[index] + period[index+1]) % n 

		if remainder == 1 && period[index+1] == 0 {
			return period[:index+1]
		}
		
		period = append(period, remainder)
	}

	return period
}

func getRemainderFibonacciNumber(f, n int) int {
	period := getPisanoPeriodNaive(n)
	return period[f%len(period)]
}

func main () {
	var f,n int

	fmt.Scanf("%v\n%v", &f, &n) 
	fmt.Printf("%v", getRemainderFibonacciNumber(f, n))
}