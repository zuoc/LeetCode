// AC 自动机
// http://blog.csdn.net/mobius_strip/article/details/22549517
// http://www.cnblogs.com/nullzx/p/7499397.html
package _30_Substring_with_Concatenation_of_All_Words

import "container/list"

type AC struct {
	root *Node

	targets []string
}

func newAC(targets []string) *AC {
	ac := &AC{
		root:    newNode('_'),
		targets: targets,
	}
	ac.buildTrie()
	ac.buildFail()
	return ac
}

func (ac *AC) buildTrie() {
	for _, target := range ac.targets {
		curr := ac.root
		for _, c := range target {
			index := c - 'a'
			if curr.childs[index] == nil {
				curr.childs[index] = newNode(c)
			}
			curr = curr.childs[index]
		}
		if curr != ac.root {
			curr.isWord = true
		}
	}
}

func (ac *AC) buildFail() {
	queue := list.New()

	for _, child := range ac.root.childs {
		child.fail = ac.root
		queue.PushBack(child)
	}

	for queue.Len() > 0 {

	}
}

func popFront(queue *list.List) *list.Element {
	ele := queue.Front()
	queue.Remove(ele)
	return ele
}

func (ac AC) search(text string) {
}

type Node struct {
	c rune

	childs []*Node

	fail *Node

	isWord bool
}

func newNode(c rune) *Node {
	return &Node{
		c:      c,
		childs: make([]*Node, 26, 26),
		fail:   nil,
		isWord: false,
	}
}
