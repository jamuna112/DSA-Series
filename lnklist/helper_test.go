package lnklist

import "fmt"

// helper function
func isEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil

}

// convert linked list to slice for better test failure output
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// build linked list from slice
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, v := range vals[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
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
