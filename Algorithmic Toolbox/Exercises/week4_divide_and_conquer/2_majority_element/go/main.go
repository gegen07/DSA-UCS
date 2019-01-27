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

func majorityElement(arr []int, left, right int) ([][]int) {
	if left + 1 == right {
		return [][]int{{arr[left]}, {}}
	}
	mid := int(left+(right - left)/2)
	leftHalf  := majorityElement(arr, left, mid)
	rightHalf := majorityElement(arr, mid, right)
	return countMerge(leftHalf, rightHalf)
}

func chunkProcess(majors, elements []int) ([]int, []int) {
	var others []int

	for _, x := range elements {
		if x == majors[0] {
			majors = append(majors, x)
		} else {
			others = append(others, x)
		}
	}

	return majors, others
}

func countMerge(leftHalf, rightHalf [][]int) [][]int  {
	
	leftMajors, rightMinors := chunkProcess(leftHalf[0], rightHalf[1])
	rightMajors, leftMinors := chunkProcess(rightHalf[0], leftHalf[1])

	if leftMajors[0] == rightMajors[0] {
		return [][]int{append(leftMajors, rightMajors...), append(leftMinors, rightMinors...)}
	} else if len(leftMajors) > len(rightMajors) {
		remaining := []int{}
		remaining = append(rightMajors, leftMinors...)
		remaining = append(remaining, rightMinors...)
		return [][]int{leftMajors, remaining}
	} else {
		remaining := []int{}
		remaining = append(leftMajors, leftMinors...)
		remaining = append(remaining, rightMinors...)
		return [][]int{rightMajors, remaining}
	}
}

func main() {

	defer writer.Flush()

	var (n int)
	scanf("%d\n", &n)
	
	array := make([]int, n)
	for index := 0; index < n; index++ {
		scanf("%d ", &array[index])
	}

	if len(majorityElement(array, 0, n)[0]) > n/2 {
		printf("1")
	} else {
		printf("0")
	}
}