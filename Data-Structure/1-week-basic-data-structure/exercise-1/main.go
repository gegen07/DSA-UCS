package main

import (
	"fmt"
	"bufio"
	"os"
	stack "github.com/gegen07/Courses/Data-Structure/1-week-basic-data-structure/exercise-1/pkg"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

var invertedBracketsCharMap = map[string]string {
	"}":"{",
	")":"(",
	"]":"[",
}

func strToStrArray(str string) []string {
	var strArray []string
	for _, char := range str {
		strArray = append(strArray, string(char))
	}

	return strArray
}

func isBalanced(str string) (int, bool) {
	stack := stack.New()
	pos := 0
	
	for _, char := range strToStrArray(str) {
		pos++
		if char == "[" || char == "(" || char == "{" {
			stack.Push(char)
		} else if char == "]" || char == ")" || char == "}" {
			if stack.Empty() {
				return pos, false
			}

			top := stack.Top()

			if top != invertedBracketsCharMap[char] {
				return pos, false
			} 
			
			stack.Pop()
		}
	}

	if stack.Size() == 1 {
		return 1, false
	}

	return pos, true
}

func main() {

	defer writer.Flush()

	var str string
	scanf("%s", &str)

	pos, ok := isBalanced(str)
	
	if ok {
		printf("Success")
	} else {
		printf("%d", pos)
	}
}