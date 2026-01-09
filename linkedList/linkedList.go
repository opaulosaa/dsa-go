package main

type LinkedList struct {
	Head *Node
}

type Node struct {
	data string //qualquer tipo de dado
	next *Node  //ponteiro do tipo Node para o proximo item da lista
}
