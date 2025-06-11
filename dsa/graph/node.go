package graph

type Node struct {
	Data     int
	Prev     *Node
	Next     *Node
	Distance int
}