package main

type Node struct {
	key  int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}
