package trie_test

import (
	"github.com/kevindragon/trie"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := trie.Create()

	trie.Add("中国")
	trie.Add("中国人民")
	trie.Add("中华人民共和国")
	trie.Add("劳动合同法")

	if trie.Find("中国人") {
		t.Error("match error")
	}
}

func BenchmarkTrieAdd(b *testing.B) {
	trie := trie.Create()

	for i := 0; i < b.N; i++ {
		trie.Add("中华人民共和国")
	}
}
