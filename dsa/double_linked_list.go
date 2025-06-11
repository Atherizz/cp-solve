package main

type DoubleNode struct {
	Data int
	Next *DoubleNode
	Prev *DoubleNode
}

type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func (list *DoubleLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *DoubleLinkedList) AddFirst(data int) {
	newNode := &DoubleNode{
		Data: data,
		Next: list.Head,
		Prev: nil,
	}
	
	if list.Head != nil {
		list.Head.Prev = newNode
	} else {
		list.Tail = newNode
	}
	
	list.Head = newNode
}

func (list *DoubleLinkedList) AddLast(data int) {
	if list.IsEmpty() {
		list.AddFirst(data)
		return
	}
	
	newNode := &DoubleNode{
		Data: data,
		Next: nil,
		Prev: list.Tail,
	}
	
	list.Tail.Next = newNode
	list.Tail = newNode
}

func (list *DoubleLinkedList) InsertAfter(data int, target int) {
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
	
	newNode := &DoubleNode{
		Data: data,
		Next: temp.Next,
		Prev: temp,
	}
	
	if temp.Next != nil {
		temp.Next.Prev = newNode
	} else {
		list.Tail = newNode
	}
	
	temp.Next = newNode
}

func (list *DoubleLinkedList) RemoveFirst() int {
	if list.IsEmpty() {
		return 0
	}
	
	data := list.Head.Data
	list.Head = list.Head.Next
	
	if list.Head != nil {
		list.Head.Prev = nil
	} else {
		list.Tail = nil
	}
	
	return data
}

func (list *DoubleLinkedList) RemoveLast() int {
	if list.IsEmpty() {
		return 0
	}
	
	if list.Head.Next == nil {
		return list.RemoveFirst()
	}
	
	data := list.Tail.Data
	list.Tail = list.Tail.Prev
	list.Tail.Next = nil
	
	return data
}

func (list *DoubleLinkedList) Remove(target int) int {
	if list.IsEmpty() {
		return 0
	}
	
	temp := list.Head
	
	if list.Head.Data == target {
		return list.RemoveFirst()
	}
	
	for temp != nil && temp.Data != target {
		temp = temp.Next
	}
	
	if temp == nil {
		return 0
	}
	
	if temp.Next != nil {
		temp.Next.Prev = temp.Prev
	} else {
		list.Tail = temp.Prev
	}
	
	temp.Prev.Next = temp.Next
	
	return temp.Data
}