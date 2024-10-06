package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCountWords(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected map[string]uint64
	}{
		{
			"Empty String Source",
			"",
			map[string]uint64{},
		},
		{
			"1 word",
			"Lorem",
			map[string]uint64{"Lorem": 1},
		},
		{
			"2 different words",
			"Lorem Ipsum",
			map[string]uint64{"Lorem": 1, "Ipsum": 1},
		},
		{
			"3 words, 2 is same",
			"Lorem Ipsum Lorem",
			map[string]uint64{"Lorem": 2, "Ipsum": 1},
		},
		{
			"3 words, 2 is same, with punctuation",
			"Lorem, Ipsum! Lorem.",
			map[string]uint64{"Lorem": 2, "Ipsum": 1},
		},
		{
			"3 similar words in different registers",
			"Lorem lorem lOrEm",
			map[string]uint64{"Lorem": 1, "lorem": 1, "lOrEm": 1},
		},
		{
			"3 words different language",
			"Lorem Ёжик Schadenfreude",
			map[string]uint64{"Lorem": 1, "Ёжик": 1, "Schadenfreude": 1},
		},
		{
			"2 words with Umlaut",
			"Ёжик Ежик",
			map[string]uint64{"Ёжик": 1, "Ежик": 1},
		},
		{
			"5 words, 2 is same, with any spaces",
			"  Lorem \t ipsum \nLorem            aaaa  bb",
			map[string]uint64{"Lorem": 2, "ipsum": 1, "aaaa": 1, "bb": 1},
		},
		{
			"Words with numbers",
			"11 L1orem \t ips43um \nLo44rem        aaaa  bb",
			map[string]uint64{"11": 1, "L1orem": 1, "ips43um": 1, "Lo44rem": 1, "aaaa": 1, "bb": 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CountWords(tc.source)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
