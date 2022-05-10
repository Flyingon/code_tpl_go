package main

import "fmt"

// Trie 树
type Trie struct {
	root *trieNode
}

type trieNode struct {
	children  [26]*trieNode
	isWordEnd bool
}

// Constructor 构造TrieTree
func Constructor() Trie {
	return Trie{
		root: &trieNode{},
	}
}

func (this *Trie) Insert(word string) {
	wordLen := len(word)
	current := this.root
	for i := 0; i < wordLen; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

func (this *Trie) Search(word string) bool {
	wordLength := len(word)
	current := this.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	wordLength := len(prefix)
	current := this.root
	for i := 0; i < wordLength; i++ {
		index := prefix[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
		if i == len(prefix)-1 {
			return true
		}
	}
	return false
}

func main() {
	trie := Constructor()
	//words := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	//for i := 0; i < len(words); i++ {
	//	trie.Insert(words[i])
	//	fmt.Println("------------------------------------------")
	//	spew.Dump(trie)
	//}
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
