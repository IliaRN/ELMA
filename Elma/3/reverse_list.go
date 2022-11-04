package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head   *Node
	length int
}

func (l *List) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l List) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *List) Reverse() {
	cursor := l.head
	var prev *Node = nil
	for cursor != nil {
		next := cursor.next
		cursor.next = prev
		prev = cursor
		cursor = next
	}
	l.head = prev
}

func main() {
	myList := List{}
	node1 := &Node{data: 12}
	node2 := &Node{data: 1}
	node3 := &Node{data: 32}
	node4 := &Node{data: 13}
	node5 := &Node{data: 31}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.printData()
	myList.Reverse()
	myList.printData()
}
