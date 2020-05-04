package leetcode

import (
	"fmt"
	"testing"
)

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
//示例:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//说明:
//
//你可以假设所有的输入都是由小写字母 a-z 构成的。
//保证所有输入均为非空字符串。

func TestTire(t *testing.T) {
	t.Run("node", func(t *testing.T) {
		n := newNode(1, false)
		fmt.Println(n.getChild(1))
		n.appendChild(newNode(1, false))
		fmt.Println(n.getChild(1))
	})
	t.Run("tire", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		fmt.Println(trie.Search("apple")) // 返回 true
		fmt.Println(trie.Search("app"))	// 返回 false
		fmt.Println(trie.StartsWith("app"))// 返回 true
		//return
		//trie.Insert("")
		trie.Insert("app")
		fmt.Println(trie.Search("app"))   // 返回 true
		fmt.Println(trie.Search("apple")) // 返回 true
		fmt.Println(trie.Search(""))      // 返回 true
	})
}

type trieNode struct {
	val   byte
	isEnd bool
	child [256]*trieNode
}
func newNode(value byte, isEnd bool) *trieNode{
	return &trieNode{value, isEnd, [256]*trieNode{}}
}

func (n *trieNode)getChild(b byte) *trieNode {
	return n.child[b]
}

func (n *trieNode)appendChild(nn *trieNode) {
	n.child[nn.val] = nn
}

func (n *trieNode)String() string {
	s := fmt.Sprintf("val:%v %v", string(n.val), n.isEnd)
	for i := 0; i < 256; i++ {
		if n.child[i] != nil {
			s += " " + n.child[i].String()
		}
	}

	return s
}
type Trie struct {
	tree *trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := newNode(0, false)
	return Trie{root}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t.tree
	for _, v := range []byte(word) {
		nn := node.getChild(v)
		if nn == nil {
			nn = newNode(v, false)
			node.appendChild(nn)
			node = nn
		}
		node = nn
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.tree
	for _, v := range []byte(word) {
		nn := node.getChild(v)
		if nn != nil {
			node = nn
		} else {
			return false
		}
	}

	if !node.isEnd {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.tree
	for _, v := range []byte(prefix) {
		nn := node.getChild(v)
		if nn != nil {
			node = nn
			continue
		}
		return false
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
