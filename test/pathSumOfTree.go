package test

/*
tree like:

		 1
	  2     3
	4  5  6  7
sum = 124+125+136+137
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	path := make([]int, 0)
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			s := 0
			path = append(path, node.Val)
			for _, n := range path {
				s *= 10
				s += n
			}
			path = path[:len(path)-1]
			sum += s
			return
		}
		path = append(path, node.Val)
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return sum
}
