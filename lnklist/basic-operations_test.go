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
