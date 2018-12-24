package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

type Coordinate struct {
	x int
	y int
}

type Signatures struct {
	points []int
}

type By func(c1, c2 *Coordinate) bool

func (by By) Sort(coordinates []Coordinate)  {
	cs := &coordinateSorter {
		coordinates: coordinates,
		by: by,
	}
	sort.Sort(cs)
}

type coordinateSorter struct {
	coordinates []Coordinate
	by    func(c1, c2 *Coordinate) bool
}

func (s *coordinateSorter) Len() int {
	return len(s.coordinates)
}

func (s *coordinateSorter) Swap(i, j int) {
	s.coordinates[i], s.coordinates[j] = s.coordinates[j], s.coordinates[i]
}

func (s *coordinateSorter) Less(i, j int) bool {
	return s.by(&s.coordinates[i], &s.coordinates[j])
}

func appendIfMissing(slice []int, i int) []int {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}

func checkElement(c Coordinate, minY int) bool {
	for j := c.x; j <= c.y; j++ {
		if j == minY {
			return true
		}
	}
	return false
}

func loopThroughAllCoordinates(c []Coordinate, minY int, result *Signatures) []Coordinate {
	for index := 0; index < len(c); index++ {
		check := checkElement(c[index], minY)

		if check {
			c = append(c[:index], c[index+1:]...)
			index--
			result.points = appendIfMissing(result.points, minY)
		}
	}

	

	return c
}

func collectSignatures(coordinates []Coordinate) Signatures {
	var result Signatures
	
	points := func(c1, c2 *Coordinate) bool {
		return c1.y < c2.y
	}

	for i := 0; i < len(coordinates); i++ {
		By(points).Sort(coordinates)
		minY := coordinates[0].y
		coordinates = loopThroughAllCoordinates(coordinates, minY, &result)
	}
	
	return result
}

func main() {

	defer writer.Flush()
	
	var n int
	scanf("%d\n", &n)

	coordinates := make([]Coordinate, n)

	for index := 0; index < n; index++ {
		scanf("%d %d\n", &coordinates[index].x, &coordinates[index].y)
	}

	result := collectSignatures(coordinates)

	printf("%v\n", len(result.points))

	for _, point := range result.points {
		printf("%v ", point)
	}
}