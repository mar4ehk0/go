package counter

import (
	"strings"
)

func CountWords(source string) map[string]uint64 {
	rawFields := strings.Fields(source)

	words := make(map[string]uint64)
	set := ",.!?"

	for _, v := range rawFields {
		word := strings.Trim(v, set)
		words[word]++
	}

	return words
}
