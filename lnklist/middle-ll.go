package lnklist

/*
	LEETCODE
	876. Middle of the Linked List

*/

func MiddleNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
