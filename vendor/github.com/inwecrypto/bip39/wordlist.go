package bip39

import (
	"strings"
	"sync"
)

// WordDictionary .
type WordDictionary struct {
	WordList       []string
	ReverseWordMap map[string]int
}

// NewWordDictionary create new bip39 word dictionary
func NewWordDictionary(words string, split string) *WordDictionary {
	wordlist := strings.Split(words, split)
	reversed := make(map[string]int)
	for i, v := range wordlist {
		reversed[v] = i
	}

	return &WordDictionary{
		WordList:       wordlist,
		ReverseWordMap: reversed,
	}
}

var mutex sync.RWMutex
var wordDictionaryZHCN = NewWordDictionary(zhCN, "\n")
var wordDictionaryENUS = NewWordDictionary(enUS, "\n")

var langs = map[string]*WordDictionary{
	"zh_CN": wordDictionaryZHCN,
	"en_US": wordDictionaryENUS,
}

// SetDict set language dictionary
func SetDict(lang string, words string, split string) {
	mutex.Lock()
	defer mutex.Unlock()

	langs[lang] = NewWordDictionary(words, split)
}

// GetDict .
func GetDict(lang string) (*WordDictionary, bool) {
	dict, ok := langs[lang]

	return dict, ok
}

// ZHCN .
func ZHCN() *WordDictionary {
	return wordDictionaryZHCN
}

// ENUS .
func ENUS() *WordDictionary {
	return wordDictionaryENUS
}
