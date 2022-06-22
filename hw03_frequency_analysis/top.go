package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	Word  string
	Count int
}

var t1 = regexp.MustCompile(`[^\wа-яА-Я-]`)

var t2 = regexp.MustCompile(`\s`)

func Top10(s string, taskWithAsteriskIsCompleted bool) []string {
	counts := make(map[string]int)

	v1 := t1.Split(s, -1)
	v2 := t2.Split(s, -1)
	if taskWithAsteriskIsCompleted {
		for _, sl := range v1 {
			if sl != "" && sl != "-" {
				counts[strings.ToLower(sl)]++
			}
		}
	} else {
		for _, sl := range v2 {
			if sl != "" {
				counts[sl]++
			}
		}
	}

	words := make([]Words, 0, len(counts))
	for k, v := range counts {
		words = append(words, Words{k, v})
	}

	sort.SliceStable(words, func(i, j int) bool {
		if words[i].Count > words[j].Count {
			return true
		}
		if words[i].Count < words[j].Count {
			return false
		}
		return words[i].Word < words[j].Word
	})

	out := make([]string, 0, 10)
	for i, w := range words {
		if i == 10 {
			break
		}
		out = append(out, w.Word)
	}
	return out
}
