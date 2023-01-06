package main

import "fmt"

func find(x int, roots []int) int {
	for x != roots[x] {
		x = roots[x]
	}
	return x
}

func union(x, y int, roots []int) bool {
	rootX := find(x, roots)
	rootY := find(y, roots)

	// check if A and B are already in the same set
	if rootX == rootY {
		return false
	}

	// merge the sets containing A and B
	roots[rootX] = rootY

	return true
}

func validTree(n int, edges [][]int) bool {
	// 1. the graph must contain n - 1 edges
	if len(edges) != n-1 {
		return false
	}

	// 2. union find
	roots := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1], roots) {
			return false
		}
	}

	return true
}

func main() {
	n := 5
	data := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	fmt.Println(validTree(n, data))
}

/*
class UnionFind:

    # For efficiency, we aren't using makeset, but instead initialising
    # all the sets at the same time in the constructor.
    def __init__(self, n):
        self.parent = [node for node in range(n)]
        # We use this to keep track of the size of each set.
        self.size = [1] * n

    # The find method, with path compression. There are ways of implementing
    # this elegantly with recursion, but the iterative version is easier for
    # most people to understand!
    def find(self, A):
        # Step 1: Find the root.
        root = A
        while root != self.parent[root]:
            root = self.parent[root]
        # Step 2: Do a second traversal, this time setting each node to point
        # directly at A as we go.
        while A != root:
            old_root = self.parent[A]
            self.parent[A] = root
            A = old_root
        return root

    # The union method, with optimization union by size. It returns True if a
    # merge happened, False if otherwise.
    def union(self, A, B):
        # Find the roots for A and B.
        root_A = self.find(A)
        root_B = self.find(B)
        # Check if A and B are already in the same set.
        if root_A == root_B:
            return False
        # We want to ensure the larger set remains the root.
        if self.size[root_A] < self.size[root_B]:
            # Make root_B the overall root.
            self.parent[root_A] = root_B
            # The size of the set rooted at B is the sum of the 2.
            self.size[root_B] += self.size[root_A]
        else:
            # Make root_A the overall root.
            self.parent[root_B] = root_A
            # The size of the set rooted at A is the sum of the 2.
            self.size[root_A] += self.size[root_B]
        return True

class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        # Condition 1: The graph must contain n - 1 edges.
        if len(edges) != n - 1: return False

        # Create a new UnionFind object with n nodes.
        unionFind = UnionFind(n)

        # Add each edge. Check if a merge happened, because if it
        # didn't, there must be a cycle.
        for A, B in edges:
            if not unionFind.union(A, B):
                return False

        # If we got this far, there's no cycles!
        return True
*/
