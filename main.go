package main

import "fmt"

func main() {

	// ********Implementação de uma Linked List************
	/*nodes := Node{
		data: "paulosa",
		next: &Node{
			data: "leles",
			next: &Node{
				data: "vava",
				next: nil,
			},
		},
	}

	fmt.Println(nodes.findLenght())*/

	nodes := &LinkedList{}
	nodes.insertAtBeginning("paulosa")
	nodes.insertAtBeginning("mariana")
	fmt.Println(nodes)

}

//------LINKED LIST-------

type LinkedList struct {
	Head *Node
}

type Node struct {
	data string //qualquer tipo de dado
	next *Node  //ponteiro do tipo Node para o proximo item da lista
}

func (ll *LinkedList) printList() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.data, " --> ")
		current = current.next
	}
	fmt.Print("null")
}

func (ll *LinkedList) findLenght() (count int) {
	//count := 0 não precisa dessa linha se ja estiver no escopo da função
	current := ll.Head
	for current != nil {
		count++
		current = current.next
	}
	return
}

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
