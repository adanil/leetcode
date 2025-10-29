package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	curr := root
	prev := make(map[*TreeNode]*TreeNode)
	used := make(map[*TreeNode]bool)
	for {
		if curr.Left != nil && !used[curr.Left] {
			prev[curr.Left] = curr
			curr = curr.Left
			continue
		}

		if !used[curr] {
			used[curr] = true
			k--
			if k == 0 {
				return curr.Val
			}
		}

		if curr.Right != nil && !used[curr.Right] {
			prev[curr.Right] = curr
			curr = curr.Right
			continue
		}

		curr = prev[curr]
	}
}
