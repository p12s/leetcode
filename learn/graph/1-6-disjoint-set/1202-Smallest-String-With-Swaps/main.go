package main

import (
	"fmt"
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	parent := make([]int, len(s))

	for idx, _ := range parent {
		parent[idx] = idx
	}

	for _, pair := range pairs {
		union(parent, pair[0], pair[1])
	}

	sets := map[int]([]int){}

	for idx, _ := range parent {
		root := find(parent, idx)
		if sets[root] == nil {
			sets[root] = []int{}
		}
		sets[root] = append(sets[root], idx)
	}
	res := make([]byte, len(s))
	for _, idxes := range sets {
		group := []byte{}
		for _, idx := range idxes {
			group = append(group, s[idx])
		}

		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})

		for idx, char := range group {
			res[idxes[idx]] = char
		}

	}

	return string(res)
}

func find(parent []int, idx int) int {
	for parent[idx] != idx {
		idx = parent[idx]
	}
	return idx
}

func union(parent []int, first int, second int) bool {
	pFirst := find(parent, first)
	pSecond := find(parent, second)

	if pFirst == pSecond {
		return false
	}
	if pFirst < pSecond {
		parent[pSecond] = pFirst
		return true
	}
	parent[pFirst] = pSecond
	return true
}

func main() {
	s := "dcab"
	pairs := [][]int{{0, 3}, {1, 2}}
	fmt.Println(smallestStringWithSwaps(s, pairs)) // bacd
}
