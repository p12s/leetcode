package main

import "fmt"

func dfs(x, y string, visited map[string]bool, g map[string]map[string]float64) float64 {
	// 1. if x or y is not in the graph return -1
	_, okX := g[x]
	_, okY := g[y]
	if !okX || !okY {
		return -1.0
	}
	// 2. if len of x edges is 0; retturn -1
	if len(g[x]) == 0 {
		return -1.0
	}
	// 3. if y in x return the result
	val, okResult := g[x][y]
	if okResult {
		return val
	}
	// 4. do dfs now, and multiply the product if answer is not -1 (which means it is not in the path of query)
	for k, _ := range g[x] {
		if !visited[k] {
			visited[k] = true
			tmp := dfs(k, y, visited, g)
			if tmp == -1.0 {
				continue
			} else {
				return tmp * g[x][k]
			}
		}
	}
	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 1. let's build the graph
	g := make(map[string]map[string]float64)
	for i, v := range equations {
		if g[v[0]] == nil {
			g[v[0]] = make(map[string]float64)
		}
		if g[v[1]] == nil {
			g[v[1]] = make(map[string]float64)
		}
		g[v[0]][v[1]] = values[i]
		g[v[1]][v[0]] = 1 / values[i]
	}

	results := make([]float64, 0, len(queries))
	// 2. answer queries
	for _, v := range queries {
		x := v[0]
		y := v[1]
		v := make(map[string]bool)
		result := dfs(x, y, v, g)
		results = append(results, result)
	}
	return results
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}

	fmt.Println(calcEquation(equations, values, queries))
}
