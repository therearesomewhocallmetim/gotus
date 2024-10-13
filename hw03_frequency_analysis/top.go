package hw03frequencyanalysis

import (
	"fmt"
	"slices"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(str string) []string {
	tokens := strings.Fields(str)
	counts := map[string]int{}

	for _, token := range tokens {
		counts[token]++
	}

	wordCounts := make([]wordCount, 0, len(counts))

	for word, count := range counts {
		wordCounts = append(wordCounts, wordCount{word, count})
	}

	slices.SortFunc(wordCounts, func(a, b wordCount) int {
		if a.count > b.count {
			return -1
		}
		if a.count < b.count {
			return 1
		}
		if a.word > b.word {
			return 1
		}
		if a.word < b.word {
			return -1
		}
		return 0
	})
	result := make([]string, 0, 10)

	for x, wc := range wordCounts {
		fmt.Printf("%s - %d\n", wc.word, wc.count)
		result = append(result, wc.word)
		if x == 9 {
			break
		}
	}
	return result
}
