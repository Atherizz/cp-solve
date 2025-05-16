package main

import "fmt"

func BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	res := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)

		for _, neighbour := range graph[node] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	return res
	
}





		// stack := stack[:len(stack)-1]


// func DFSRecursive(graph map[int][]int, node int, visited map[int]bool, result *[]int) {

// 	if visited[node] {
// 		return
// 	}

// 	visited[node] = true
// 	*result = append(*result, node)

// 	for _, neighbour := range graph[node] {
// 		if !visited[neighbour] {
// 			DFS(graph, neighbour, visited, result)
// 		}
// 	}

// }

func main() {
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {7},
		6: {},
		7: {},
	}
	fmt.Println(BFS(graph, 1))
	fmt.Println(DFS(graph,1))

	// visited := make(map[int]bool)
	// result := []int{}
	

	// DFSRecursive(graph, 1, visited, &result)
	// fmt.Println("DFS Recursive Result:", result)

}