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

// push back
func (ll *LinkedList) PushBack(data int) {

	newNode := &Node{Data: data}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	//traverse to the last node
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	//append the new node
	current.Next = newNode
}

func (l *LinkedList) PrintList() string {

	var result string
	temp := l.Head

	for temp != nil {
		result += fmt.Sprintf("%v->", temp.Data)
		temp = temp.Next
	}

	return result
}
