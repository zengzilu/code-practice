package _24

import "math"

type TreeNode struct {
	Var   int
	Left  *TreeNode
	Right *TreeNode
}

// 包含节点与节点子孙节点的路径和最大值
var curMax int

func maxPathSum(root *TreeNode) int {
	curMax = math.MinInt

	_ = maxSubNode(root)
	return curMax
}

func maxSubNode(root *TreeNode) int {
	// 叶子节点的 最大子孙轮径和等于自己
	if root.Left == nil && root.Right == nil {
		if root.Var > curMax {
			curMax = root.Var
		}
		return root.Var
	} else {
		leftChild, rightChild := math.MinInt, math.MinInt
		if root.Left != nil {
			leftChild = maxSubNode(root.Left)
		}
		if root.Right != nil {
			rightChild = maxSubNode(root.Right)
		}

		// 最大路径和不出当前根节点的可能情况，记录最大值
		if leftChild > 0 && rightChild > 0 {
			if leftChild+rightChild+root.Var > curMax {
				curMax = leftChild + rightChild + root.Var
			}
		}

		// 计算最大路径和包含根节点外部节点的情况
		rootMax := root.Var
		maxChild := max(leftChild, rightChild)
		if maxChild > 0 {
			rootMax += maxChild
		}
		if rootMax > curMax {
			curMax = rootMax
		}

		return rootMax
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
