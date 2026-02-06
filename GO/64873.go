package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func getTopWords(wordMap map[string]int, n int) []string {
	type kv struct {
		word  string
		count int
	}
	var pairs []kv
	for word, count := range wordMap {
		pairs = append(pairs, kv{word, count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
	top := make([]string, 0, n)
	for i := 0; i < n && i < len(pairs); i++ {
		top = append(top, pairs[i].word)
	}
	return top
}

func cleanWord(word string) string {
	var b strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func AnalyzeText(text string) {
	re := regexp.MustCompile(`[ \t\n\r\f\v]+`)
	wordsRaw := re.Split(text, -1)
	var words []string
	for _, w := range wordsRaw {
		cleaned := cleanWord(w)
		if cleaned != "" {
			words = append(words, cleaned)
		}
	}
	totalWords := len(words)
	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w]++
	}
	uniqueWords := len(wordCount)
	var maxWord string
	maxCount := 0
	for w, c := range wordCount {
		if c > maxCount {
			maxWord = w
			maxCount = c
		}
	}
	top5 := getTopWords(wordCount, 5)
	fmt.Printf("Количество слов: %d\n", totalWords)
	fmt.Printf("Количество уникальных слов: %d\n", uniqueWords)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", maxWord, maxCount)
	fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
	for _, w := range top5 {
		fmt.Printf("\"%s\": %d раз\n", w, wordCount[w])
	}
}

func main() {
	AnalyzeText("one two two three three three four four four four five five five five five")
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	AnalyzeText("Я так люблю море. Я на море. Я так люблю плавать. Море! Я море!!! ЛЮБЛЮ МОРЕ")
}
