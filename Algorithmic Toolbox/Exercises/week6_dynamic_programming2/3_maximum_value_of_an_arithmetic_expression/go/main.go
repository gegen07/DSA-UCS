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

func min(x ... int) int {
	minValue := x[0]

	for _, value := range x {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func switchOperation(a, b int, op rune) int {
	switch op {
		case '+':
			return a+b
		case '-':
			return a-b
		case '*':
			return a*b
		default:
			panic("Operation not recognized!")
	} 
}

func separateOpRune(expression []rune) ([]int, []rune) {
	var ( digits []int
	 	ops []rune
	)

	for _, value := range expression {
		if value >= 48 && value <= 57 { //ASCII table
			digits = append(digits, (int(value)-48)) // 48 dec number is equal to 0
		} else {
			ops = append(ops, value)
		}
	}

	return digits, ops
}

func getMinMaxValue(i, j int, dpMin, dpMax [][]int, ops []rune) (int, int) {
	minValue := +100000
	maxValue := -100000 
	for k := i; k < j; k++ {
		a := switchOperation(dpMax[i][k], dpMax[k+1][j], ops[k])
		b := switchOperation(dpMax[i][k], dpMin[k+1][j], ops[k])
		c := switchOperation(dpMin[i][k], dpMax[k+1][j], ops[k])
		d := switchOperation(dpMin[i][k], dpMin[k+1][j], ops[k])

		maxValue = max(maxValue,a,b,c,d)
		minValue = min(minValue,a,b,c,d)

	}

	return minValue, maxValue
}

func setMatrixDP(length int) ([][]int) {
	dpMatrix := make([][]int, length)

	for i := range dpMatrix {
		dpMatrix[i] = make([]int, length)
	}

	return dpMatrix
}

func placingParenthesis(expression []rune) int {
	digits, ops := separateOpRune(expression)
	dpMax := setMatrixDP(len(digits))
	dpMin := setMatrixDP(len(digits))

	for i := 0; i < len(digits); i++ {
		dpMax[i][i] = digits[i]
		dpMin[i][i] = digits[i]
	}

	for s := 0; s < len(digits); s++ {
		for i := 0; i < (len(digits) - s -1); i++ {
			j := i + s + 1
			min, max := getMinMaxValue(i, j, dpMin, dpMax, ops)
			dpMax[i][j] = max
			dpMin[i][j] = min
		}
	}

	return dpMax[0][len(digits)-1]
}

func main() {

	defer writer.Flush()
	
	var exp string

	scanf("%s\n", &exp)
	result := placingParenthesis([]rune(exp))
	printf("\n%v", result)
}