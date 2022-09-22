package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := []string{"apel", "apel", "mangga", "jambu", "jambu", "jambu", "jambu", "apel", "belimbing", "durian"}
	var countfruit map[string]int = map[string]int{}

	sortfruit := make([]string, 0, len(countfruit))

	for i := 0; i < len(fruits); i++ {
		if _, ok := countfruit[fruits[i]]; ok {
			countfruit[fruits[i]] += 1
		} else {
			countfruit[fruits[i]] = 1
		}
	}

	for k := range countfruit {
		sortfruit = append(sortfruit, k)
	}
	sort.Sort(sort.StringSlice(sortfruit)) //sort by Key
	sort.SliceStable(sortfruit, func(i, j int) bool {
		return countfruit[sortfruit[i]] > countfruit[sortfruit[j]] //sort by value source from GeeksforGeeks
	})

	for _, k := range sortfruit {
		fmt.Println(k, countfruit[k])
	}
}
