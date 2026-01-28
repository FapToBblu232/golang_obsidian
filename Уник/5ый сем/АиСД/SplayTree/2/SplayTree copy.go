package main

import (
	"bufio"
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

func (tree *Tree) rightRotate(y *Node) {
	x := y.left
	if x == nil {
		return
	}
	y.left = x.right
	if x.right != nil {
		x.right.father = y
	}
	x.father = y.father
	if y.father == nil {
		tree.root = x
	} else if y.father.left == y {
		y.father.left = x
	} else {
		y.father.right = x
	}
	x.right = y
	y.father = x
}

func (tree *Tree) leftRotate(x *Node) {
	y := x.right
	if y == nil {
		return
	}
	x.right = y.left
	if y.left != nil {
		y.left.father = x
	}
	y.father = x.father
	if x.father == nil {
		tree.root = y
	} else if x.father.left == x {
		x.father.left = y
	} else {
		x.father.right = y
	}
	y.left = x
	x.father = y
}

func (tree *Tree) splay(x *Node) {
	for x.father != nil {
		p := x.father
		g := p.father
		if g == nil { // один zig
			if x == p.left {
				tree.rightRotate(p)
			} else {
				tree.leftRotate(p)
			}
		} else if x == p.left && p == g.left { // zig zig левый-левый
			tree.rightRotate(g)
			tree.rightRotate(p)
		} else if x == p.right && p == g.right { // zig-zig дважды правый
			tree.leftRotate(g)
			tree.leftRotate(p)
		} else if x == p.right && p == g.left {
			tree.leftRotate(p)
			tree.rightRotate(g)
		} else {
			tree.rightRotate(p)
			tree.leftRotate(g)
		}
	}
}

func (tree *Tree) search(key int) *Node {
	if tree.root == nil {
		return nil
	}
	cur := tree.root
	var lastNotNil *Node
	for cur != nil {
		lastNotNil = cur
		if key == cur.key {
			tree.splay(cur)
			return cur
		}
		if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if lastNotNil != nil {
		tree.splay(lastNotNil)
	}
	return nil
}

func (tree *Tree) add(key int, value string, writer *bufio.Writer) {
	if tree.root == nil {
		tree.root = &Node{key: key, value: value}
		return
	}
	cur := tree.root
	var lastNotNil *Node
	for cur != nil {
		lastNotNil = cur
		if key == cur.key {
			writer.WriteString("error\n")
			tree.splay(cur)
			return
		}
		if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	newNode := &Node{key: key, value: value, father: lastNotNil}
	if key < lastNotNil.key {
		lastNotNil.left = newNode
	} else {
		lastNotNil.right = newNode
	}
	tree.splay(newNode)
}

func (tree *Tree) set(key int, value string) bool {
	node := tree.search(key)
	if node == nil {
		return false
	}
	node.value = value
	return true
}

func (tree *Tree) delete(key int) bool {
	node := tree.search(key)
	if node == nil {
		return false
	}

	leftSub := node.left
	rightSub := node.right

	if leftSub != nil {
		leftSub.father = nil
	}
	if rightSub != nil {
		rightSub.father = nil
	}

	if leftSub == nil {
		tree.root = rightSub
	} else {
		maxLeft := leftSub
		for maxLeft.right != nil {
			maxLeft = maxLeft.right
		}
		tree.splay(maxLeft)
		maxLeft.right = rightSub
		if rightSub != nil {
			rightSub.father = maxLeft
		}
		tree.root = maxLeft
	}
	return true
}

func (tree *Tree) min() *Node {
	if tree.root == nil {
		return nil
	}
	cur := tree.root
	for cur.left != nil {
		cur = cur.left
	}
	tree.splay(cur)
	return cur
}

func (tree *Tree) max() *Node {
	if tree.root == nil {
		return nil
	}
	cur := tree.root
	for cur.right != nil {
		cur = cur.right
	}
	tree.splay(cur)
	return cur
}

func (tree *Tree) print(writer *bufio.Writer) {
	if tree.root == nil {
		writer.WriteString("_\n")
		return
	}

	queue := []*Node{tree.root}

	for len(queue) > 0 {
		nextLevel := []*Node{}
		lineEmpty := true

		for i, node := range queue {
			if i > 0 {
				writer.WriteByte(' ')
			}
			if node == nil {
				writer.WriteByte('_')
				nextLevel = append(nextLevel, nil, nil)
			} else {
				writer.WriteString(fmt.Sprintf("[%d %s", node.key, node.value))
				if node.father != nil {
					writer.WriteString(fmt.Sprintf(" %d]", node.father.key))
				} else {
					writer.WriteByte(']')
				}
				nextLevel = append(nextLevel, node.left, node.right)
				if node.left != nil || node.right != nil {
					lineEmpty = false
				}
			}
		}
		writer.WriteByte('\n')

		if lineEmpty {
			break
		}
		queue = nextLevel
	}
}

func main() {
	tree := &Tree{}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "add":
			if len(fields) != 3 {
				writer.WriteString("error\n")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil {
				writer.WriteString("error\n")
				continue
			}
			tree.add(key, fields[2], writer)

		case "set":
			if len(fields) != 3 {
				writer.WriteString("error\n")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil || !tree.set(key, fields[2]) {
				writer.WriteString("error\n")
				continue
			}

		case "delete":
			if len(fields) != 2 {
				writer.WriteString("error\n")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil || !tree.delete(key) {
				writer.WriteString("error\n")
				continue
			}

		case "search":
			if len(fields) != 2 {
				writer.WriteString("error\n")
				continue
			}
			key, err := strconv.Atoi(fields[1])
			if err != nil {
				writer.WriteString("error\n")
				continue
			}
			temp := tree.search(key)
			if temp == nil {
				writer.WriteString("0\n")
			} else {
				writer.WriteString(fmt.Sprintf("1 %s\n", temp.value))
			}

		case "max":
			if len(fields) != 1 {
				writer.WriteString("error\n")
				continue
			}

			temp := tree.max()
			if temp == nil {
				writer.WriteString("error\n")
				continue
			}
			writer.WriteString(fmt.Sprintf("%d %s\n", temp.key, temp.value))

		case "min":
			if len(fields) != 1 {
				writer.WriteString("error\n")
				continue
			}

			temp := tree.min()
			if temp == nil {
				writer.WriteString("error\n")
				continue
			}
			writer.WriteString(fmt.Sprintf("%d %s\n", temp.key, temp.value))

		case "print":
			tree.print(writer)
		default:
			writer.WriteString("error\n")
		}
	}
}
