package main

import "fmt"

func (ll *LinkedList) printList() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.data, " --> ")
		current = current.next
	}
	fmt.Print("null")
}
