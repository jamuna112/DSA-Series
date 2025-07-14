package main

import (
	"dsa-series/lnklist"
	"fmt"
)

func main() {

	list := &lnklist.DoublyList{}

	list.Push_Front(5)
	list.Push_Front(6)
	list.Push_Front(7)
	fmt.Println("===== Push front function =====")
	list.PrintDoublyList()

	list.Push_Back(1)
	fmt.Println("\n===== Push back function =====")
	list.PrintDoublyList()

	fmt.Println("\n===== Pop Front function =====")
	list.Pop_Front()
	list.PrintDoublyList()

	fmt.Println("\n===== Pop Back function =====")
	list.Pop_Back()
	list.PrintDoublyList()

}
