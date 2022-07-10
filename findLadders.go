package main

import "fmt"

func run126() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	s := findLadders(beginWord, endWord, wordList)
	fmt.Println(s)
	// [hit hot dot dog log cog cog]
	// [hit hot dot dog log cog]
	// [hit hot dot dog lot log cog]
	// [hit hot dot lot log dog cog]
	// [hit hot dot lot log dog cog]
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// minCount := int(^uint(0) >> 1)
	paths := make([][]string, 0)
	path := make([]string, 0)
	wordMap := findLadders_buildNeighborMap(wordList)
	path = append(path, beginWord)
	begins := findLadders_getNeighborWord(beginWord, wordList)
	// count := 0
	var dfs func([]string, []string)
	dfs = func(begins []string, path []string) {
		for _, word := range begins {
			// count++
			if findLadders_contains(word, path) {
				continue
			}
			path = append(path, word)
			if word == endWord {
				paths = append(paths, path)
				break
			}
			nexts, _ := wordMap[word]
			dfs(nexts, path)
		}
	}
	dfs(begins, path)

	return paths
}

func findLadders_contains(word string, wordList []string) bool {
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}

func findLadders_buildNeighborMap(wordList []string) map[string][]string {
	wordMap := make(map[string][]string, len(wordList))
	for i := 0; i < len(wordList); i++ {
		list := append([]string{}, wordList[:i]...)
		list = append(list, wordList[i+1:]...)
		wordMap[wordList[i]] = findLadders_getNeighborWord(
			wordList[i], list)
	}
	return wordMap
}

func findLadders_getNeighborWord(word string, wordList []string) []string {
	strs := make([]string, 0)
	for _, w := range wordList {
		if len(w) != len(word) {
			continue
		}
		sword := word[:]
		sw := w[:]
		diffs := 0
		for i := 0; i < len(sword); i++ {
			if sw[i] != sword[i] {
				diffs++
				if diffs > 1 {
					break
				}
			}
		}
		if diffs == 1 {
			strs = append(strs, w)
		}
	}
	return strs
}
