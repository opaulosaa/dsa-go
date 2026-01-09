package main

func (ll *LinkedList) findLenght() (count int) {
	//count := 0 não precisa dessa linha se ja estiver no escopo da função
	current := ll.Head
	for current != nil {
		count++
		current = current.next
	}
	return
}
