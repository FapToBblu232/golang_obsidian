package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func contains(vertexes []string, target string) bool {
	for i, _ := range vertexes {
		if target == vertexes[i] {
			return true
		}
	}
	return false
}

func insertOr(graph map[string][]string, pair []string) {
	v1, v2 := pair[0], pair[1]
	if !contains(graph[v1], v2) {
		graph[v1] = append(graph[v1], v2)
	}
}

func insertNotOr(graph map[string][]string, pair []string) {
	v1, v2 := pair[0], pair[1]
	if !contains(graph[v1], v2) {
		graph[v1] = append(graph[v1], v2)
	}
	if !contains(graph[v2], v1) {
		graph[v2] = append(graph[v2], v1)
	}
}

func dfs(graph map[string][]string, start string, visited map[string]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println(start)

	neighbors := append([]string{}, graph[start]...)
	sort.Strings(neighbors)
	for _, vert := range neighbors {
		dfs(graph, vert, visited)
	}
}

func bfs(graph map[string][]string, start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}
		visited[node] = true
		fmt.Println(node)

		neighbors := append([]string{}, graph[node]...)
		sort.Strings(neighbors)
		for _, vert := range neighbors {
			if !visited[vert] {
				queue = append(queue, vert)
			}
		}
	}
}

func main() {
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	first_line := scanner.Text()
	fields := strings.Fields(first_line)
	graphType := fields[0]
	startVertex := fields[1]
	traversalType := fields[2]

	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			continue
		}
		pair := strings.Fields(str)
		if graphType == "d" {
			insertOr(graph, pair)
		} else if graphType == "u" {
			insertNotOr(graph, pair)
		}
	}
	if traversalType == "b" {
		bfs(graph, startVertex)
	} else if traversalType == "d" {
		visited := make(map[string]bool)
		dfs(graph, startVertex, visited)
	}
}
