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
	key    int
	value  string
	left   *Node
	right  *Node
	father *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) search(key int) *Node {
	if tree.root == nil {
		return nil
	}
	return tree.root.search(key)
}

func (node *Node) search(key int) *Node {
	if node == nil {
		return nil
	}
	if key == node.key {
		return node
	}
	if key > node.key {
		return node.right.search(key)
	}
	return node.left.search(key)
}

func (tree *Tree) add(key int, value string) error {
	if tree.root == nil {
		tree.root = &Node{key: key, value: value}
		return nil
	}
	return tree.root.add(key, value)
}

func (node *Node) add(key int, value string) error {
	if key == node.key {
		return errors.New("error")
	} else if key < node.key {
		if node.left == nil {
			node.left = &Node{key: key, value: value, father: node}
			return nil
		}
		return node.left.add(key, value)
	} else {
		if node.right == nil {
			node.right = &Node{key: key, value: value, father: node}
			return nil
		}
		return node.right.add(key, value)
	}
}

func (tree *Tree) set(key int, value string) error {
	temp := tree.search(key)
	if temp == nil {
		return errors.New("error")
	}
	temp.value = value
	return nil
}

func (tree *Tree) delete(key int) error {
	temp := tree.root.search(key)
	if temp == nil {
		return errors.New("error")
	}
	tree.root = tree.root.delete(key)
	return nil
}

func (node *Node) delete(key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = node.left.delete(key)
	} else if key > node.key {
		node.right = node.right.delete(key)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.right == nil {
			node.left.father = node.father
			return node.left
		} else if node.left == nil {
			node.right.father = node.father
			return node.right
		} else {
			temp := node.left.max() // нашли то, что макс слева (У него не будет ребёнка справа)
			node.key = temp.key
			node.value = temp.value
			node.left = node.left.delete(node.key)
		}
	}
	return node // если всё хорршо удалилось, то это мы воткнём в нужные нам места
}

func (tree *Tree) print() {
	if tree.root == nil {
		fmt.Println("_")
		return
	}

	queue := []*Node{tree.root}
	nextLevel := []*Node{}

	fmt.Printf("[%d %s]\n", tree.root.key, tree.root.value)

	for len(queue) > 0 {
		for _, node := range queue {
			if node == nil {
				nextLevel = append(nextLevel, nil, nil)
				continue
			}
			nextLevel = append(nextLevel, node.left, node.right)
		}

		noOne := true
		for _, n := range nextLevel {
			if n != nil {
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
			if n == nil {
				fmt.Print("_")
			} else {
				fmt.Printf("[%d %s %d]", n.key, n.value, n.father.key)
			}
		}
		fmt.Println()

		queue = nextLevel
		nextLevel = []*Node{}
	}
}

func (tree *Tree) min() *Node {
	if tree.root == nil {
		return nil
	}
	return tree.root.min()
}

func (node *Node) min() *Node {
	if node.left == nil {
		return node
	}
	return node.left.min()
}

func (tree *Tree) max() *Node {
	if tree.root == nil {
		return nil
	}
	return tree.root.max()
}

func (node *Node) max() *Node {
	if node.right == nil {
		return node
	}
	return node.right.max()
}

func main() {
	tree := &Tree{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, " ") || strings.HasSuffix(line, " ") || strings.Contains(line, "  ") {
			fmt.Println("error")
			continue
		}
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "add":
			if len(fields) != 3 {
				fmt.Println("error")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("error")
				continue
			}
			err = tree.add(key, fields[2])
			if err != nil {
				fmt.Println("error")
				continue
			}

		case "set":
			if len(fields) != 3 {
				fmt.Println("error")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("error")
				continue
			}
			err = tree.set(key, fields[2])
			if err != nil {
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
			err = tree.delete(key)
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
			temp := tree.search(key)
			if temp == nil {
				fmt.Println("0")
			} else {
				fmt.Printf("1 %s\n", temp.value)
			}

		case "max":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}

			temp := tree.max()
			if temp == nil {
				fmt.Println("error")
				continue
			}
			fmt.Printf("%d %s\n", temp.key, temp.value)

		case "min":
			if len(fields) != 1 {
				fmt.Println("error")
				continue
			}

			temp := tree.min()
			if temp == nil {
				fmt.Println("error")
				continue
			}
			fmt.Printf("%d %s\n", temp.key, temp.value)

		case "print":
			tree.print()
		default:
			fmt.Println("error")
		}
	}
}
