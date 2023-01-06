package main

import "fmt"

func findCircleNum(edges [][]int) int {
	roots := make([]int, len(edges))
	for i := 0; i < len(edges[0]); i++ {
		roots[i] = i
	}

	start := 0
	for i := 0; i < len(edges); i++ {
		for j := start; j < len(edges[0]); j++ {
			if edges[i][j] != 0 {
				union(i, j, roots)
			}
		}
		start++
	}

	m := make(map[int]int)
	for i := 0; i < len(roots); i++ {
		m[roots[i]] = 1
	}
	return len(m)
}

func find(x int, roots []int) int {
	return roots[x]
}

func union(x, y int, roots []int) {
	rootX := find(x, roots)
	rootY := find(y, roots)
	if rootX != rootY {
		for i := 0; i < len(roots); i++ {
			if roots[i] == rootY {
				roots[i] = rootX
			}
		}
	}
}

func main() {
	edges := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	res := findCircleNum(edges)
	fmt.Println(res)
}
