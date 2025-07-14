package lnklist

import "fmt"

type Node1 struct {
	data int
	next *Node1
	prev *Node1
}

type DoublyList struct {
	head *Node1
	tail *Node1
}

func (dll *DoublyList) Push_Front(value int) {

	newNode := &Node1{data: value, next: nil, prev: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *DoublyList) Push_Back(value int) {
	newNode := &Node1{data: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func (dll *DoublyList) Pop_Front() *Node1 {

	if dll.head == nil {
		return nil
	}

	temp := dll.head
	dll.head = dll.head.next

	if dll.head != nil {
		dll.head.prev = nil
	} else {
		dll.tail = nil
	}

	temp.next = nil
	return temp
}

func (dll *DoublyList) Pop_Back() *Node1 {

	if dll.head == nil {
		return nil
	}

	temp := dll.tail
	dll.tail = dll.tail.prev

	if dll.tail != nil {
		dll.tail.next = nil
	}

	temp.next = nil
	return temp
}

func (dll *DoublyList) PrintDoublyList() {
	temp := dll.head

	for temp != nil {
		fmt.Printf("%d <=>", temp.data)
		temp = temp.next
	}

}
