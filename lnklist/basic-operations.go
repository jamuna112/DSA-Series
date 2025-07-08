package lnklist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct { 
	Head *Node
}

func (l *LinkedList) PushFront(data int) {
	newNode := &Node{Data: data}
	current := l.Head
	if current == nil {
		l.Head = newNode
	} else {
		newNode.Next = current
		l.Head = newNode
	}

}

func (l *LinkedList) PrintList() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()	
}