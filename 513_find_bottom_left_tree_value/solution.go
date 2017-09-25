// BFS
package _513_find_bottom_left_tree_value

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		ele := popFront(queue)

		if node, ok := ele.Value.(*TreeNode); ok {
			root = node

			if root.Right != nil {
				queue.PushBack(root.Right)
			}

			if root.Left != nil {
				queue.PushBack(root.Left)
			}
		}

	}
	return root.Val

}

func popFront(queue *list.List) *list.Element {
	ele := queue.Front()
	queue.Remove(ele)
	return ele
}
