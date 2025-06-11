package graph

import "fmt"

func NewNode(data int, distance int) *Node {
	return &Node{
		Data:     data,
		Distance: distance,
		Next:     nil,
		Prev:     nil,
	}
}

type DoubleLinkedList struct {
	Head *Node
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		Head: nil,
	}
}

func (dll *DoubleLinkedList) IsEmpty() bool {
	return dll.Head == nil
}

func (dll *DoubleLinkedList) AddFirst(data int, distance int) {
	newNode := NewNode(data, distance)
	if dll.IsEmpty() {
		dll.Head = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
}

func (dll *DoubleLinkedList) RemoveFirst() {
	if dll.IsEmpty() {
		fmt.Println("List kosong, tidak bisa dihapus")
		return
	}

	if dll.Head.Next == nil {
		dll.Head = nil
		return
	}

	dll.Head = dll.Head.Next
	dll.Head.Prev = nil
}

func (dll *DoubleLinkedList) GetDistance(index int) int {
	current := dll.Head

	for i := 0; i < index; i++ {
		if current == nil {
			fmt.Println("Index melebihi panjang list")
			return -1 // atau panic/log error
		}
		current = current.Next
	}

	if current == nil {
		fmt.Println("Index melebihi panjang list")
		return -1
	}

	return current.Distance
}

func (dll *DoubleLinkedList) RemoveLast() {
	if dll.IsEmpty() {
		fmt.Println("List kosong, tidak bisa dihapus")
		return
	}

	current := dll.Head

	for current.Next != nil {
		current = current.Next
	}

	if current == dll.Head {
		dll.Head = nil
		return
	}

	current.Prev.Next = nil
}

func(dll *DoubleLinkedList) GetSize() int {

	if dll.IsEmpty() {
		return 0
	}
	
	counter := 0
	current := dll.Head

	for current != nil {
		current = current.Next
		counter++
	}

	return counter
}

func (dll *DoubleLinkedList) Remove(index int) {
	if dll.IsEmpty() {
		fmt.Println("List kosong, tidak bisa dihapus")
		return
	}

	if index < 0 {
		fmt.Println("Index tidak valid")
		return
	}

	if index == 0 {
		dll.RemoveFirst()
		return
	}
	temp := dll.Head

	for i := 0; i < index; i++ {
		if temp == nil {
			fmt.Println("Index melebihi panjang list")
			return
		}
		temp = temp.Next
	}

	if temp.Next == nil {
		dll.RemoveLast()
		return
	}

	temp.Next.Prev = temp.Prev
	temp.Prev.Next = temp.Next
}
