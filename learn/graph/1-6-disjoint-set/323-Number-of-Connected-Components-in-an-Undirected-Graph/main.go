package main

import "fmt"

func countComponents(n int, edges [][]int) int {
	roots := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
	}

	find := func(x int) int {
		return roots[x]
	}

	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			for i := 0; i < len(roots); i++ {
				if roots[i] == rootY {
					roots[i] = rootX
				}
			}
		}
	}
	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	m := make(map[int]int)
	for _, root := range roots {
		m[root] = 1
	}

	return len(m)
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}

	fmt.Println(countComponents(n, edges))

}
