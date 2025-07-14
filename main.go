package main

import "dsa-series/lnklist"

func main() {

	list := &lnklist.DoublyList{}

	list.Push_Front(5)
	list.Push_Front(6)
	list.Push_Front(7)
	list.PrintDoublyList()
}
