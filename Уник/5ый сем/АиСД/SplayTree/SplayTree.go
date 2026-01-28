package main

import (
	"bufio"
	"fmt"
	"os"

	"runtime"
	"strconv"
	"strings"
	// "time"
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

func (tree *Tree) zig(x *Node) {
	p := x.father
	if p.father == nil {
		tree.root = x
		x.father = nil
	} else {
		g := p.father
		if p == g.left {
			g.left = x
		} else {
			g.right = x
		}
		x.father = g
	}
	if x == p.left {
		p.left = x.right
		if x.right != nil {
			x.right.father = p
		}
		x.right = p
	} else {
		p.right = x.left
		if x.left != nil {
			x.left.father = p
		}
		x.left = p
	}
	p.father = x
}

func (tree *Tree) zig_zig(x *Node) {
	tree.zig(x.father)
	tree.zig(x)
}

func (tree *Tree) zig_zag(x *Node) {
	tree.zig(x)
	tree.zig(x)
}

func (tree *Tree) splay(x *Node) {
	for x.father != nil {
		p := x.father
		if p.father == nil {
			tree.zig(x)
		} else if (x == p.left) == (p == p.father.left) {
			tree.zig_zig(x)
		} else {
			tree.zig_zag(x)
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

func (tree *Tree) add(key int, value string) {
	if tree.root == nil {
		tree.root = &Node{key: key, value: value}
		return
	}
	cur := tree.root
	var lastNotNil *Node
	for cur != nil {
		lastNotNil = cur
		if key == cur.key {
			fmt.Println("error")
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

func (tree *Tree) set(key int, value string) {
	node := tree.search(key)
	if node == nil {
		fmt.Println("error")
		return
	}
	node.value = value
}

func (tree *Tree) delete(key int) {
	node := tree.search(key)
	if node == nil {
		fmt.Println("error")
		return
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

func (tree *Tree) print() {
	if tree.root == nil {
		fmt.Println("_")
		return
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "[%d %s]\n", tree.root.key, tree.root.value)

	queue := []any{}
	if tree.root.left != nil {
		queue = append(queue, tree.root.left)
	} else {
		queue = append(queue, 1)
	}

	if tree.root.right != nil {
		queue = append(queue, tree.root.right)
	} else {
		switch queue[len(queue)-1].(type) {
		case int:
			return
		case *Node:
			queue = append(queue, 1)
		}
	}

	anyOne := true
	for anyOne {
		anyOne = false
		next := []any{}
		for i := 0; i < len(queue)-1; i++ {
			switch v := queue[i].(type) {
			case int:
				fmt.Fprint(&sb, strings.Repeat("_ ", v))
				if len(next) == 0 {
					next = append(next, 2*v)
				} else {
					switch z := next[len(next)-1].(type) {
					case *Node:
						if len(next) != 0 {
							next = append(next, 2*v)
						}
					case int:
						if len(next) != 0 {
							next[len(next)-1] = z + 2*v
						}
					}
				}
			case *Node:
				fmt.Fprintf(&sb, "[%d %s %d] ", v.key, v.value, v.father.key)
				if v.left != nil {
					next = append(next, v.left)
					anyOne = true
				} else {
					if len(next) == 0 {
						next = append(next, 1)
					} else {
						switch z := next[len(next)-1].(type) {
						case int:
							next[len(next)-1] = z + 1
						case *Node:
							if len(next) != 0 {
								next = append(next, 1)
							}
						}
					}
				}
				if v.right != nil {
					next = append(next, v.right)
					anyOne = true
				} else {
					if len(next) == 0 {
						next = append(next, 1)
					} else {
						switch z := next[len(next)-1].(type) {
						case int:
							next[len(next)-1] = z + 1
						case *Node:
							if len(next) != 0 {
								next = append(next, 1)
							}
						}
					}
				}
			}
		}
		// вывод последнего
		temp := queue[len(queue)-1]
		switch v := temp.(type) {
		case int:
			fmt.Fprint(&sb, strings.Repeat("_ ", v-1))
			fmt.Fprint(&sb, "_\n")
			if len(next) == 0 {
				next = append(next, 2*v)
			} else {
				switch z := next[len(next)-1].(type) {
				case *Node:
					if len(next) != 0 {
						next = append(next, 2*v)
					}
				case int:
					if len(next) != 0 {
						next[len(next)-1] = z + 2*v
					}
				}
			}
		case *Node:
			fmt.Fprintf(&sb, "[%d %s %d]\n", v.key, v.value, v.father.key)
			if v.left != nil {
				next = append(next, v.left)
				anyOne = true
			} else {
				if len(next) == 0 {
					next = append(next, 1)
				} else {
					switch z := next[len(next)-1].(type) {
					case int:
						next[len(next)-1] = z + 1
					case *Node:
						if len(next) != 0 {
							next = append(next, 1)
						}
					}
				}
			}
			if v.right != nil {
				next = append(next, v.right)
				anyOne = true
			} else {
				if len(next) == 0 {
					next = append(next, 1)
				} else {
					switch z := next[len(next)-1].(type) {
					case int:
						next[len(next)-1] = z + 1
					case *Node:
						if len(next) != 0 {
							next = append(next, 1)
						}
					}
				}
			}
		}
		queue = next
	}
	fmt.Print(sb.String())
}

func main() {
	tree := &Tree{}
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
				if err != nil {
					fmt.Println("error")
					continue
				}
				tree.add(key, fields[2])
			} else if len(fields) == 2 {
				key, err := strconv.Atoi(fields[1])
				if err != nil {
					fmt.Println("error")
					continue
				}
				tree.add(key, "")
			} else {
				fmt.Println("error")
				continue
			}

		case "set":
			if len(fields) == 3 {
				key, err := strconv.Atoi(fields[1])
				if err != nil {
					fmt.Println("error")
					continue
				}
				tree.set(key, fields[2])
			} else if len(fields) == 2 {
				key, err := strconv.Atoi(fields[1])
				if err != nil {
					fmt.Println("error")
					continue
				}
				tree.set(key, "")
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
			tree.delete(key)

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
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys = %v KB\n", m.Sys/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
