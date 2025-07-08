package main

import (
	"dsa-series/lnklist"
)

func main() {
	var ll lnklist.LinkedList
	ll.PushFront(10)
	ll.PushFront(20)
	ll.PushFront(30)
	ll.PrintList()
}
