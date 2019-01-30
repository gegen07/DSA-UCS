package main

import (
	"fmt"
	"bufio"
	"os"
	stackMax "github.com/gegen07/Courses/Data-Structure/1-week-basic-data-structure/exercise-4/stackmax"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {

	defer writer.Flush()

	stk := stackMax.New()

	var (
		quantity      int
		value         int
		maxValueArray []int
		userCommand   string
	)

	scanf("%d\n", &quantity)

	for i := 0; i < quantity; i++ {
		writer.Flush()
		scanf("%s\n", &userCommand)

		switch userCommand {
		case "push":
			scanf("%d\n", &value)
			stk.Push(value)
		case "pop":
			stk.Pop()
		case "max":
			maxValue, err := stk.Max()

			if !err {
				maxValueArray = append(maxValueArray, maxValue)
			}
		} 
	}

	if len(maxValueArray) < 1 {
		printf("\n")
	} else {
		for _, x := range maxValueArray {
			printf("%d\n", x)
		}
	}
}