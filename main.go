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
