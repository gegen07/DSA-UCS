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

type Loot struct {
	weight         int
	value          int
	valuePerWeight int
}

type By func(l1, l2 *Loot) bool

func (by By) Sort(loots []Loot) {
	ls := &lootSorter {
		loots: loots,
		by:    by,
	}
	sort.Sort(ls)
}

type lootSorter struct {
	loots []Loot
	by    func(l1, l2 *Loot) bool
}

func (s *lootSorter) Len() int {
	return len(s.loots)
}

func (s *lootSorter) Swap(i, j int) {
	s.loots[i], s.loots[j] = s.loots[j], s.loots[i]
}

func (s *lootSorter) Less(i, j int) bool {
	return s.by(&s.loots[i], &s.loots[j])
}

func getMaximumValueLoot(loots []Loot, W, n int) float64 {
	valuePerWeight := func(l1, l2 *Loot) bool {
		return l1.valuePerWeight > l2.valuePerWeight
	}

	By(valuePerWeight).Sort(loots)

	curWeight := 0
	finalValue := 0.0

	for index := 0; index < n; index++ {
		if (curWeight + loots[index].weight <= W) {
			curWeight += loots[index].weight
			finalValue += float64(loots[index].value)
		} else {
			remain := W - curWeight
			finalValue += float64(loots[index].value) * (float64(remain) / float64(loots[index].weight))
			break
		}
	}

	return finalValue
}

func main() {

	defer writer.Flush()

	var n, W int
	scanf("%d %d\n", &n, &W)

	var loots = make([]Loot, n)

	for index := 0; index < n; index++ {
		scanf("%d %d\n", &loots[index].value, &loots[index].weight)
		loots[index].valuePerWeight = loots[index].value / loots[index].weight
	}

	printf("%.2f", getMaximumValueLoot(loots, W, n))
}