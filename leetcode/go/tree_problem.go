package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
}

func maxDepth(root *TreeNode) int {
	return height(root)
}

	func isSameTree(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}


func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
