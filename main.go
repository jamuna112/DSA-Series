package main

import (
	"dsa-series/lnklist"
	"dsa-series/stack"
	"fmt"
)

func main() {

	list := &lnklist.DoublyList{}

	list.Push_Front(5)
	list.Push_Front(6)
	list.Push_Front(7)
	// fmt.Println("===== Push front function =====")
	// list.PrintDoublyList()

	list.Push_Back(1)
	// fmt.Println("\n===== Push back function =====")
	// list.PrintDoublyList()

	// fmt.Println("\n===== Pop Front function =====")
	// list.Pop_Front()
	// list.PrintDoublyList()

	// fmt.Println("\n===== Pop Back function =====")
	// list.Pop_Back()
	// list.PrintDoublyList()

	//nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	// ans := stack.NextGreaterElement1(nums1, nums2)
	// // fmt.Print("ans is ", ans)

	valStack := stack.ReverseStack(nums2)
	for _, val := range valStack {
		fmt.Printf("%d ", val)
	}

}
