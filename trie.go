// Trie tree
package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	value    string
	children map[string]*TrieNode
	isEnd    bool
}

func Create() *Trie {
	root := &Trie{
		&TrieNode{
			value:    "",
			children: make(map[string]*TrieNode),
		},
	}

	return root
}

func (tree *Trie) Add(str string) {
	add(tree.root, str)
}

func (tree *Trie) Find(str string) bool {
	return find(tree.root, str)
}

func (tree *Trie) Travel() []string {
	var words []string

	var travel2 func(*TrieNode, string)
	travel2 = func(node *TrieNode, pfx string) {
		curValue := node.value

		str := pfx
		str += curValue
		pfx = str

		if node.isEnd {
			words = append(words, str)
			str = ""
		}

		if len(node.children) > 0 {
			for _, childNode := range node.children {
				travel2(childNode, pfx)
			}
		}
	}
	travel2(tree.root, "")

	return words
}

func find(node *TrieNode, str string) bool {
	if len(str) == 0 {
		return false
	}

	chars := []rune(str)
	firstChar := string(chars[0])
	lastChars := string(chars[1:])

	if childNode, ok := node.children[firstChar]; ok {
		if len(lastChars) == 0 {
			if childNode.isEnd {
				return true
			} else {
				return false
			}
		} else {
			return find(childNode, lastChars)
		}
	} else {
		return false
	}
}

func add(node *TrieNode, str string) {
	chars := []rune(str)

	if len(chars) == 0 {
		node.isEnd = true
	} else {
		firstChar := string(chars[0])
		lastChars := string(chars[1:])

		if childNode, ok := node.children[firstChar]; ok {
			add(childNode, lastChars)
		} else {
			newNode := &TrieNode{
				value:    firstChar,
				children: make(map[string]*TrieNode),
			}
			node.children[firstChar] = newNode
			add(newNode, lastChars)
		}
	}
}
