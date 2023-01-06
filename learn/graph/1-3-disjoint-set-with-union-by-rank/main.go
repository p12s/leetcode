package main

import "fmt"

func find(x int) int { // O(N)
	for x != roots[x] {
		x = roots[x]
	}
	return x
}

func union(x, y int) { // O(N)
	rootX := find(x)
	rootY := find(y)
	if rootX != rootY {
		if rank[rootX] > rank[rootY] {
			roots[rootY] = rootX
		} else if rank[rootX] < rank[rootY] {
			roots[rootX] = rootY
		} else {
			roots[rootY] = rootX
			rank[rootX] += 1
		}
	}
}

func connected(x, y int) bool { // O(N)
	return find(x) == find(y)
}

var (
	size  = 10
	roots = make([]int, size)
	rank  = make([]int, size)
)

func main() {

	// 1-2-5-6-7 3-8-9 4
	for i := 0; i < size; i++ { // O(N)
		roots[i] = i
		rank[i] = 1
	}

	union(1, 2)
	union(2, 5)
	union(5, 6)
	union(6, 7)
	union(3, 8)
	union(8, 9)

	fmt.Println(connected(1, 5)) // true
	fmt.Println(connected(5, 7)) // true
	fmt.Println(connected(4, 9)) // false

	// 1-2-5-6-7 3-8-9-4
	union(9, 4)
	fmt.Println(connected(4, 9)) // true
}
