package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type Words struct {
	Word  string
	Count int
}

func Top10(s string) []string {
	counts := make(map[string]int)
	t := regexp.MustCompile(`\s+`)
	v := t.Split(s, -1)
	for _, sl := range v {
		if sl != " " {
			counts[sl]++
		}
	}
	words := []Words{}
	for k, v := range counts {
		words = append(words, Words{k, v})
	}
	sort.SliceStable(words, func(i, j int) bool { return words[i].Word < words[j].Word })
	sort.SliceStable(words, func(i, j int) bool { return words[i].Count > words[j].Count })
	var out []string
	for i, w := range words {
		if w.Count > 1 {
			out = append(out, w.Word)
		}
		if i == 9 {
			break
		}
	}
	return out
}
