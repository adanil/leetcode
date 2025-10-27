package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	prev := make(map[*TreeNode]*TreeNode)
	prev[root] = nil
	curr := root
	result := make([]int, 0)
	used := make(map[*TreeNode]bool)
	for {
		if curr.Left != nil && !used[curr.Left] {
			prev[curr.Left] = curr
			curr = curr.Left
			continue
		}

		if !used[curr] {
			result = append(result, curr.Val)
			used[curr] = true
		}

		if curr.Right != nil && !used[curr.Right] {
			prev[curr.Right] = curr
			curr = curr.Right
			continue
		}

		curr = prev[curr]
		if curr == nil {
			break
		}
	}
	return result
}

func main() {

}
