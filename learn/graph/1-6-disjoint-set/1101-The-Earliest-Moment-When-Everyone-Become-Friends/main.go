package main

import (
	"fmt"
	"sort"
)

func find(person int, group []int) int {
	if group[person] != person {
		group[person] = find(group[person], group)
	}
	return group[person]
}

func union(a, b int, group, rank []int) bool {
	groupA := find(a, group)
	groupB := find(b, group)

	var isMerged bool
	if groupA == groupB {
		return isMerged
	}

	isMerged = true
	if rank[groupA] > rank[groupB] {
		group[groupB] = groupA
	} else if rank[groupA] < rank[groupB] {
		group[groupA] = groupB
	} else {
		group[groupA] = groupB
		rank[groupB] += 1
	}

	return isMerged
}

func earliestAcq(logs [][]int, n int) int {
	if len(logs) < n-1 {
		return -1
	}

	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	group := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		group[i] = i
		rank[i] = 0
	}

	// initially, we treat each individual as a separate group
	groupCnt := n

	// we merge the groups along the way.
	for _, log := range logs {
		timestamp := log[0]
		friendA := log[1]
		friendB := log[2]

		if union(friendA, friendB, group, rank) {
			groupCnt -= 1
		}

		// the moment when all individuals are connected to each other.
		if groupCnt == 1 {
			return timestamp
		}
	}

	// there are still more than one groups left, i.e. not everyone is connected
	return -1
}

func main() {
	n := 4
	logs := [][]int{{4, 1, 2}, {7, 3, 1}, {0, 2, 0}, {3, 0, 3}, {1, 0, 1}}
	fmt.Println(earliestAcq(logs, n)) // 3

	n2 := 6
	logs2 := [][]int{
		{20190101, 0, 1},
		{20190104, 3, 4},
		{20190107, 2, 3},
		{20190211, 1, 5},
		{20190224, 2, 4},
		{20190301, 0, 3},
		{20190312, 1, 2},
		{20190322, 4, 5},
	}
	fmt.Println(earliestAcq(logs2, n2)) // 20190301
}
