package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type BstTree struct {
	Root *TreeNode
}

func (t *BstTree) IsEmpty() bool {
	return t.Root == nil
}


func (t *BstTree) Find(node *TreeNode) bool {
	return t.FindHelper(t.Root, node)
}

func (t *BstTree) FindHelper(root *TreeNode, node *TreeNode) bool {
	if root == nil {
		return false
	}

	if node.Data == root.Data {
		return true
	} else if node.Data < root.Data {
		return t.FindHelper(root.Left, node)
	} else {
		return t.FindHelper(root.Right, node)
	}
}

func (t *BstTree) Insert(data int) {
	newNode := &TreeNode{
		Data: data,
	}
	t.Root = t.InsertHelper(t.Root, newNode)
}

func (t *BstTree) InsertHelper(root *TreeNode, node *TreeNode) *TreeNode {
	data := node.Data
	if root == nil {
		root = node
		return root
	}

	if data < root.Data {
		root.Left = t.InsertHelper(root.Left, node)
	} else if data > root.Data {
		root.Right = t.InsertHelper(root.Right, node)
	}

	return root
}

func (t *BstTree) Successsor(node *TreeNode) int {
	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	return successor.Data
}

func (t *BstTree) Predecessor(node *TreeNode) int {
	predecessor := node.Left
	for predecessor.Right != nil {
		predecessor = predecessor.Right
	}

	return predecessor.Data
}

func (t *BstTree) Delete(target int) {
	t.Root = t.DeleteHelper(t.Root, target)
}

func (t *BstTree) DeleteHelper(node *TreeNode, target int) *TreeNode {
	if t.IsEmpty() {
		return nil
	}

	if node == nil {
		return nil
	}

	if target < node.Data {
		node.Left = t.DeleteHelper(node.Left, target)
	} else if target > node.Data {
		node.Right = t.DeleteHelper(node.Right, target)
	} else {
		if node.Left == nil && node.Right == nil {
			node = nil
		} else if node.Right != nil {
			node.Data = t.Successsor(node)
			node.Right = t.DeleteHelper(node.Right, node.Data)
		} else if node.Left != nil {
			node.Data = t.Predecessor(node)
			node.Left = t.DeleteHelper(node.Left, node.Data)
		}
	}

	return node
}

func (t *BstTree) Display() {
	if t.IsEmpty() {
		return
	}
	t.preOrder(t.Root)
	fmt.Println()
}

func (t *BstTree) preOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Print(node.Data, " ")
	t.preOrder(node.Left)
	t.preOrder(node.Right)
}

func (t *BstTree) inOrder(node *TreeNode) {
	if node == nil {
		return
	}

	t.inOrder(node.Left)
	t.inOrder(node.Right)
}

func (t *BstTree) PostOrder(node *TreeNode) {
	if node == nil {
		return
	}

	t.PostOrder(node.Left)
	t.PostOrder(node.Right)
	fmt.Print(node.Data, " ")
}

func main() {
	tree := BstTree{
		Root: nil,
	}

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(80)
	tree.Insert(60)

	tree.Display()

	tree.Delete(50)

	tree.Display()
}
