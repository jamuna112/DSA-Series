package lnklist

/**
LEETCODE: 206 Reverse Linked List

 * Definition for singly-linked list.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	var prev, curr, next *ListNode

	curr = head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// func PrintList(head *ListNode) {
// 	for head != nil {
// 		fmt.Printf("%d ", head.Val)
// 		head = head.Next
// 	}
// 	fmt.Println()
// }
