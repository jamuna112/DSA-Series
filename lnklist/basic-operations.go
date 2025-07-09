package lnklist

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

// Pop front
func (l *LinkedList) PopFront() (int, bool) {

	if l.Head == nil {
		return 0, false
	}

	val := l.Head.Data
	l.Head = l.Head.Next

	return val, true

}

// pop back
func (l *LinkedList) PopBack() (int, bool) {
	if l.Head == nil {
		return 0, false // list is empty
	}

	//if list has only one node
	if l.Head.Next == nil {
		temp := l.Head.Data
		l.Head = nil
		return temp, true
	}

	//if list has multiple nodes, traverse to the second last node
	prev := l.Head
	curr := l.Head.Next

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil //remove last node
	return curr.Data, true
}
