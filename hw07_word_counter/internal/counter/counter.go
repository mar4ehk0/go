package counter

import (
	"strings"
)

func CountWords(source string) map[string]uint64 {
	rawFields := strings.Fields(source)

	words := make(map[string]uint64)

	for _, v := range rawFields {
		words[v]++
	}

	return words
}
