package lnklist

import (
	"testing"
)

func TestPushFront(t *testing.T) {

	tests := []struct {
		name string
		data int
	}{
		{"test 1", 10},
		{"test 2", 20},
		{"test 3", 30},
		{"test 4", 40},
		{"test 5", 50},
	}

	for _, tt := range tests {
		t.Run((tt.name), func(t *testing.T) {
			ll := &LinkedList{}

			ll.PushFront(tt.data)

			if ll.Head == nil && ll.Head.Data != tt.data {
				t.Errorf("Expected head to be %d, got nil", tt.data)
			} else if ll.Head != nil && ll.Head.Data != tt.data {
				t.Errorf("Expected head to be %d, got %d", tt.data, ll.Head.Data)
			}

		})
	}

}

func TestPushBack(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected string
	}{
		{"Empty list", []int{}, ""},
		{"Single element", []int{10}, "10->"},
		{"Multiple elements", []int{1, 2, 3, 4}, "1->2->3->4->"},
		{"Duplicates elements", []int{2, 2}, "2->2->"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}

			for _, pb := range tt.values {
				ll.PushBack(pb)
			}

			got := ll.PrintList()
			if got != tt.expected {
				t.Errorf("PushBack() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestPopFront(t *testing.T) {
	list := &LinkedList{}

	//test empty list
	val, okay := list.PopFront()

	if okay || val != 0 {
		t.Errorf("test failed, expected 0, and false empty list but got %d, %v", val, okay)
	}

	//add items
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	list.PushBack(40)

	//pop first item
	val, ok := list.PopFront()

	if val != 10 || !ok {
		t.Errorf("expected %d and %v but got %v, and %v", 10, true, val, ok)
	}

	//pop second item

	val, ok = list.PopFront()
	if val != 20 || ok != true {
		t.Errorf("expected %d and %v but got %v, and %v", 20, true, val, ok)

	}
}

func TestPopBack(t *testing.T) {
	list := &LinkedList{}

	//when list is empty
	val, okay := list.PopBack()

	if val != 0 || okay {
		t.Errorf("test failed, expected 0, and false on empty list but got %d and %v", val, okay)
	}

	//add one item to the list
	list.PushFront(100)

	val, okay = list.PopBack()

	if val != 100 || okay != true {
		t.Errorf("test failed, expected 100, and true from list but got %d and %v", val, okay)

	}

	//add more items to the list
	list.PushFront(200)
	list.PushFront(300)
	list.PushFront(400)

	val, okay = list.PopBack()

	if val != 200 || okay != true {
		t.Errorf("test failed, expected 200, and true from list but got %d and %v", val, okay)

	}

}

func TestPrintList(t *testing.T) {

	tests := []struct {
		name     string
		list     *LinkedList
		expected string
	}{
		{"Empty list", &LinkedList{Head: nil}, ""},
		{"Single node", &LinkedList{Head: &Node{Data: 10}}, "10->"},
		{"Multiple nodes", &LinkedList{Head: &Node{Data: 1, Next: &Node{Data: 2, Next: &Node{Data: 3}}}}, "1->2->3->"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if output := tt.list.PrintList(); output != tt.expected {
				t.Errorf("PrintList() = %q, want %q", output, tt.expected)
			}
		})
	}
}
