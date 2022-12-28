package main

// You are given a directed graph consisting of N vertices, numbered from 1 to N, and N edges.
// The graph is described by two arrays, A and B, both of length N. A pair A[K], B[K] (for K from 0 to N-1)
// describes a directed edge from vertex A[K] to vertex B[K]. Note that the graph might contain self-loops (edges where A[K] = B[K])
// and/or multiple edges between the same pair of vertices.
// Your task is to check whether the given graph is a cycle. A graph is a cycle if it is possible to start
// at some vertex and, by following the provided edges, visit all the other vertices and return to the starting point.
//
// write a function that, given two arrays A and B of N integers each, describing a directed graph, returns
// if the graph is a cycle and false otherwise

func IsCycle(A, B []int) bool {
	// Create an adjacency list representation of the graph
	adjList := make(map[int][]int)
	for i := 0; i < len(A); i++ {
		adjList[A[i]] = append(adjList[A[i]], B[i])
	}

	// Use a depth-first search to check if the graph is a cycle
	visited := make(map[int]bool)
	for i := 1; i <= len(A); i++ {
		if !visited[i] && dfs(i, i, adjList, visited) {
			return true
		}
	}

	return false
}

func dfs(u, start int, adjList map[int][]int, visited map[int]bool) bool {
	visited[u] = true
	for _, v := range adjList[u] {
		if !visited[v] {
			if dfs(v, start, adjList, visited) {
				return true
			}
		} else if v == start {
			return true
		}
	}
	return false
}
