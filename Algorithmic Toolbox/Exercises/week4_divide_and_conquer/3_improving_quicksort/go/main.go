package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return min + rand.Int()%((max-min)+1)
}
func partition3(arr *[]int, low, high int) (int, int) {
	pivot := (*arr)[low]
	i := low
	pivot1 := low
	pivot2 := high

	for i <= pivot2 {
		if (*arr)[i] < pivot {
			(*arr)[pivot1], (*arr)[i] = (*arr)[i], (*arr)[pivot1]
			pivot1++
			i++
		} else if (*arr)[i] > pivot {
			(*arr)[i], (*arr)[pivot2] = (*arr)[pivot2], (*arr)[i]
			pivot2--
		} else {
			i++
		}
	}

	return pivot1, pivot2
}

func randomizedPartition(arr *[]int, low, high int) (int, int) {
	k := random(low, high)
	(*arr)[k], (*arr)[low] = (*arr)[low], (*arr)[k]
	return partition3(arr, low, high)
}

func randomizedQuickSort(arr *[]int, low, high int) {
	if low < high {
		m1, m2 := randomizedPartition(arr, low, high)
		randomizedQuickSort(arr, low, m1-1)
		randomizedQuickSort(arr, m2+1, high)
	}
	return
}

func main() {

	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	array := make([]int, n)

	for index := 0; index < n; index++ {
		scanf("%d ", &array[index])
	}

	randomizedQuickSort(&array, 0, n-1)

	for index := 0; index < n; index++ {
		printf("%d ", array[index])
	}
}
