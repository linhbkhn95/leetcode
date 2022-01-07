package main

import (
	"fmt"
	"sort"
)

type wordDictInfo struct {
	value   string
	lenWord int
}

func wordBreak(s string, wordDict []string) bool {
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})
	flash := make(map[string]interface{})
	prefixWordDictMapping := make(map[byte][]wordDictInfo)
	for _, word := range wordDict {
		l := len(word)
		words, ok := prefixWordDictMapping[word[0]]
		if !ok {
			prefixWordDictMapping[word[0]] = []wordDictInfo{
				{lenWord: l, value: word},
			}
		} else {
			prefixWordDictMapping[word[0]] = append(words, wordDictInfo{lenWord: l, value: word})
		}
	}

	return backtracking(s, len(s), prefixWordDictMapping, prefixWordDictMapping[s[0]], flash)
}

func backtracking(s string, l int, preprefixWordDictMapping map[byte][]wordDictInfo, eligibleWords []wordDictInfo, flash map[string]interface{}) bool {
	if len(eligibleWords) == 0 {
		return false
	}
	if _, ok := flash[s]; ok {
		return false
	}
	for _, word := range eligibleWords {
		if word.lenWord > l || s[:word.lenWord] != word.value {
			continue
		}
		newS := s[word.lenWord:]
		if newS == "" {
			return true
		}

		if backtracking(newS, l-word.lenWord, preprefixWordDictMapping, preprefixWordDictMapping[newS[0]], flash) {
			return true
		}
		flash[newS] = nil
	}
	return false
}

func main() {
	trie := NewTrie()

	trie.Insert("linhtd")
	trie.Insert("xuna")
	fmt.Println(trie.Find("linhtd"))
	fmt.Println(trie.Find("xuan"))
}

type TrieNode struct {
	tries  [26]*TrieNode
	isWord bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.tries[index] == nil {
			current.tries[index] = &TrieNode{}
		}
		current = current.tries[index]
	}
	current.isWord = true
}

func (t *Trie) Find(word string) bool {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.tries[index] == nil {
			return false
		}
		current = current.tries[index]
	}
	return current.isWord
}

// func (t *Trie) Extract(prefix string) bool {
// 	current := t.root
// 	for i := 0; i < len(word); i++ {
// 		index := word[i] - 'a'
// 		if current.tries[index] == nil {
// 			return false
// 		}
// 		current = current.tries[index]
// 	}
// 	return current.isWord
// }
