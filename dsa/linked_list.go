package main

type Node struct {
	Data int
	Next *Node
}

type SingleLinkedList struct {
	Head *Node
}

func (list *SingleLinkedList) IsEmpty() bool {
	return list.Head == nil
}


func (list *SingleLinkedList) AddFirst(data int) {
	newNode := &Node{
		Data: data,
		Next: list.Head,
	}
	list.Head = newNode
}

func (list *SingleLinkedList) AddLast(data int) {

	if list.IsEmpty() {
		list.AddFirst(data)
		return
	}
	
	temp := list.Head
	
	for temp.Next != nil {
		temp = temp.Next
	}

	newNode := &Node{
		Data: data,
		Next: nil,
	}

	temp.Next = newNode
}

func (list *SingleLinkedList) InsertAfter(data int, target int) {
	if list.IsEmpty() {
	return
	}

	temp := list.Head

	for temp != nil && temp.Data != target {
		temp = temp.Next
	} 

	if temp == nil {
		return
	}

	newNode := &Node{
	Data: data,
	Next: temp.Next,
	}

	temp.Next = newNode	
	}

func (list *SingleLinkedList) RemoveFirst() int {
	if list.IsEmpty() {
	return 0
	}

	data := list.Head
	list.Head = list.Head.Next

	return data.Data
}

func (list *SingleLinkedList) RemoveLast() int {
	 
	var prev *Node 

	if list.IsEmpty() {
	return 0
	}

	temp := list.Head

	if list.Head.Next == nil {
		return list.RemoveFirst()
	}

	for temp.Next != nil {
		prev = temp
		temp = temp.Next
	}

	prev.Next = nil

	return temp.Data

}

func (list *SingleLinkedList) Remove(target int) int {
	
	if list.IsEmpty() {
		return 0
	}
	
	var prev *Node 
	temp := list.Head

	if list.Head.Data == target {
		return list.RemoveFirst()
	}

	for temp != nil && temp.Data != target {
		prev = temp
		temp = temp.Next
	}

	if temp == nil {
		return 0
	}

	prev.Next = temp.Next

	return temp.Data
}






