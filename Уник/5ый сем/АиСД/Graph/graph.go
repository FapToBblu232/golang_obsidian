package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Queue struct {
	data     []string
	len      int
	capacity int
	first    int
	last     int
}

func (queue *Queue) resize() {
	newCap := queue.capacity * 2
	newData := make([]string, newCap)

	for i := 0; i < queue.len; i++ {
		newData[i] = queue.data[(queue.first+i)%queue.capacity]
	}

	queue.data = newData
	queue.capacity = newCap
	queue.first = 0
	queue.last = queue.len
}

func (queue *Queue) pushb(input string) {
	if queue.len == queue.capacity {
		queue.resize()
	}
	queue.data[queue.last] = input
	queue.last = (queue.last + 1) % queue.capacity
	queue.len++
}

func (queue *Queue) popf() string {
	if queue.len == 0 {
		return ""
	}
	answ := queue.data[queue.first]
	queue.first = (queue.first + 1) % queue.capacity
	queue.len--
	return answ
}

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

	// Это то же самое, что и код вида
	var neighbors []string
	neighbors = append(neighbors, graph[start]...)
	// neighbors := append([]string{}, graph[start]...)
	sort.Strings(neighbors)
	for _, vert := range neighbors {
		dfs(graph, vert, visited)
	}
}

func bfs(graph map[string][]string, start string) {
	visited := make(map[string]bool)
	// кстати, разве при операции queue = queue[1:] мы не просто создаём новый слайс,
	// у которого адрес нулевого элемента указывает на след элемент изначального?
	// т.е. массив же в памяти не менялся
	queue := Queue{
		data:     make([]string, 4),
		len:      0,
		capacity: 4,
		first:    0,
		last:     0,
	}
	queue.pushb(start)

	for queue.len > 0 {
		node := queue.popf()

		if visited[node] {
			continue
		}
		visited[node] = true
		fmt.Println(node)

		// Это то же самое, что и код вида
		var neighbors []string
		neighbors = append(neighbors, graph[node]...)
		// neighbors := append([]string{}, graph[node]...)

		sort.Strings(neighbors)
		for _, vert := range neighbors {
			if !visited[vert] {
				queue.pushb(vert)
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
