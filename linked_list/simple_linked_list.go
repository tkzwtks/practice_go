package main

import "fmt"

type listNode struct {
	data int
	next *listNode
}

func listLength(head *listNode) int {
	currentNode := head
	count := 0

	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}
	return count
}

func unshift(head *listNode, node *listNode) {
	node.next = head
}

func push(last *listNode, node *listNode) {
	last.next = node
}

func traverse(startNode *listNode) {
	current := startNode
	count := 0
	for current != nil {
		count++
		fmt.Printf("%d: number: %d\n", count, current.data)
		current = current.next
	}
	fmt.Printf("end\n")
}

func main() {
	node1 := &listNode{1, nil}
	node2 := &listNode{2, node1}
	node3 := &listNode{3, node2}
	node4 := &listNode{4, node3}
	node5 := &listNode{5, node4}

	traverse(node5)

	fmt.Println("===== unshift =====")
	newHead := &listNode{99, nil}
	unshift(node5, newHead)
	traverse(newHead)

	fmt.Println("===== push =====")
	tail := &listNode{0, nil}
	push(node1, tail)
	traverse(newHead)
}
