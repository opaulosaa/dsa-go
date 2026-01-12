package main

func (ll *LinkedList) insertAtBeginning(data string) {
	newNode := &Node{data: data, next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) insertAtEnd(data string) {
	newNode := &Node{data: data, next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}
