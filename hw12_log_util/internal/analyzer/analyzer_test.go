package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAnalyze(t *testing.T) {
	stat := NewStat("GET")
	expected := 2

	stat.Analyze("LOREM ipsum GET test")
	stat.Analyze("GET LOREM ipsum test")

	actual := stat.Value

	assert.Equal(t, expected, actual)
}
