// 字典树
package _208_implement_trie

type Node struct {
	c rune

	childrens []*Node

	isLeaf bool
}

func NewNode(c rune) *Node {
	return &Node{
		c:         c,
		childrens: make([]*Node, 26, 26),
		isLeaf:    false,
	}
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: NewNode('_'),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		index := c - 'a'
		if node.childrens[index] == nil {
			node.childrens[index] = NewNode(c)
		}
		node = node.childrens[index]
	}
	node.isLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		index := c - 'a'
		if node.childrens[index] == nil {
			return false
		}
		node = node.childrens[index]
	}
	return node.isLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		index := c - 'a'
		if node.childrens[index] == nil {
			return false
		}
		node = node.childrens[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
