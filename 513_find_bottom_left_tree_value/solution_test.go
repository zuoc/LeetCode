package _513_find_bottom_left_tree_value

import (
	"testing"
	"fmt"
)

func Test_findBottomLeftValue(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}

	fmt.Println(findBottomLeftValue(root))
}
