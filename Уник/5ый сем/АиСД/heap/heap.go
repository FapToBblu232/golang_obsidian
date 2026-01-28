package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   int
	value string
}

type Heap struct {
	arr   []*Node
	size  int
	exist map[int]int // чтобы search(key) был за O(1), а не O(n)
	// Тут мы храним ключ + его индекс в массиве
}

func (heap *Heap) siftDown(i int) {
	for {
		left := i*2 + 1
		right := i*2 + 2
		great := i

		if (left < heap.size) && (heap.arr[left].key < heap.arr[great].key) {
			great = left
		}
		if (right < heap.size) && (heap.arr[right].key < heap.arr[great].key) {
			great = right
		}
		if great == i {
			break
		}
		heap.arr[i], heap.arr[great] = heap.arr[great], heap.arr[i]
		heap.exist[heap.arr[i].key] = i
		heap.exist[heap.arr[great].key] = great
		i = great
	}
}

func (heap *Heap) siftUp(i int) {
	for i > 0 {
		parent := int((i - 1) / 2)
		if heap.arr[i].key > heap.arr[parent].key {
			return
		}
		heap.arr[i], heap.arr[parent] = heap.arr[parent], heap.arr[i]
		heap.exist[heap.arr[i].key] = i
		heap.exist[heap.arr[parent].key] = parent
		i = parent
	}
}

func (heap *Heap) add(key int, value string) bool {
	if _, ok := heap.exist[key]; ok {
		return false
	}
	heap.arr = append(heap.arr, &Node{key: key, value: value})
	heap.size++
	heap.exist[key] = heap.size - 1
	heap.siftUp(heap.size - 1)
	return true
}

func (heap *Heap) extract() *Node {
	if heap.size <= 0 {
		return nil
	}
	answ := heap.arr[0]
	heap.arr[0] = heap.arr[heap.size-1]
	heap.exist[heap.arr[0].key] = 0
	heap.size--
	heap.arr = heap.arr[:heap.size]
	delete(heap.exist, answ.key)
	if heap.size != 0 {
		heap.siftDown(0)
	}
	return answ
}

func (heap *Heap) delete(key int) error {
	ind, ok := heap.exist[key]
	if !ok {
		return errors.New("error")
	}

	lastInd := heap.size - 1
	lastKey := heap.arr[lastInd].key
	delete(heap.exist, key)
	if ind != lastInd {
		heap.arr[ind], heap.arr[lastInd] = heap.arr[lastInd], heap.arr[ind]
		heap.exist[lastKey] = ind
	}
	heap.size--
	heap.arr = heap.arr[:heap.size]
	if ind < heap.size {
		heap.siftUp(ind)
		heap.siftDown(ind)
	}
	return nil
}

func (heap *Heap) set(key int, value string) bool {
	if _, ok := heap.exist[key]; !ok {
		return false
	}
	heap.arr[heap.exist[key]].value = value
	return true
}

func (heap *Heap) search(key int) int {
	if _, ok := heap.exist[key]; !ok {
		return -1
	}
	return heap.exist[key]
}

func (heap *Heap) min() int {
	if heap.size <= 0 {
		return -1
	}
	return 0
}

func (heap *Heap) max() int {
	if heap.size <= 0 {
		return -1
	}
	maxi := len(heap.arr) - 1
	for i := int(heap.size / 2); i < heap.size; i++ {
		if heap.arr[i].key > heap.arr[maxi].key {
			maxi = i
		}
	}
	return maxi
}

func (heap *Heap) print() {
	if heap.size <= 0 {
		fmt.Println("_")
		return
	}

	queue := []int{0}
	nextLevel := []int{}

	fmt.Printf("[%d %s]\n", heap.arr[0].key, heap.arr[0].value)

	for len(queue) > 0 {
		for _, ind := range queue {
			if ind == -1 {
				nextLevel = append(nextLevel, -1, -1)
				continue
			}
			if (ind*2)+1 < heap.size {
				nextLevel = append(nextLevel, (ind*2)+1)
			} else {
				nextLevel = append(nextLevel, -1)
			}
			if (ind*2)+2 < heap.size {
				nextLevel = append(nextLevel, (ind*2)+2)
			} else {
				nextLevel = append(nextLevel, -1)
			}
		}

		noOne := true
		for _, n := range nextLevel {
			if n != -1 {
				noOne = false
				break
			}
		}
		if noOne {
			return
		}

		for i, n := range nextLevel {
			if i > 0 {
				fmt.Print(" ")
			}
			if n == -1 {
				fmt.Print("_")
			} else {
				fmt.Printf("[%d %s %d]", heap.arr[n].key, heap.arr[n].value, heap.arr[(n-1)/2].key)
			}
		}
		fmt.Println()

		queue = nextLevel
		nextLevel = []int{}
	}
}

func main() {
	heap := &Heap{exist: map[int]int{}}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "add":
			if len(fields) == 3 {
				key, err := strconv.Atoi(fields[1])
				if err != nil || !heap.add(key, fields[2]) {
					fmt.Println("error")
					continue
				}
			} else if len(fields) == 2 {
				key, err := strconv.Atoi(fields[1])
				if err != nil || !heap.add(key, "") {
					fmt.Println("error")
					continue
				}
			} else {
				fmt.Println("error")
				continue
			}

		case "set":
			if len(fields) == 3 {
				key, err := strconv.Atoi(fields[1])
				if err != nil || !heap.set(key, fields[2]) {
					fmt.Println("error")
					continue
				}
			} else if len(fields) == 2 {
				key, err := strconv.Atoi(fields[1])
				if err != nil || !heap.set(key, "") {
					fmt.Println("error")
					continue
				}
			} else {
				fmt.Println("error")
				continue
			}

		case "delete":
			if len(fields) != 2 {
				fmt.Println("error")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("error")
				continue
			}
			err = heap.delete(key)
			if err != nil {
				fmt.Println("error")
				continue
			}

		case "search":
			if len(fields) != 2 {
				fmt.Println("error")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("error")
				continue
			}
			temp := heap.search(key)
			if temp == -1 {
				fmt.Println("0")
			} else {
				fmt.Printf("1 %d %s\n", temp, heap.arr[temp].value)
			}

		case "max":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}

			temp := heap.max()
			if temp == -1 {
				fmt.Println("error")
				continue
			}
			fmt.Printf("%d %d %s\n", heap.arr[temp].key, temp, heap.arr[temp].value)

		case "min":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}

			temp := heap.min()
			if temp == -1 {
				fmt.Println("error")
				continue
			}
			fmt.Printf("%d %d %s\n", heap.arr[temp].key, temp, heap.arr[temp].value)

		case "print":
			heap.print()

		case "extract":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}
			node := heap.extract()
			if node != nil {
				fmt.Printf("%d %s\n", node.key, node.value)
			} else {
				fmt.Println("error")
				continue
			}
		default:
			fmt.Println("error")
		}
	}
}
