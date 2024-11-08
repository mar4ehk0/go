package analyzer

import (
	"fmt"
	"strings"
)

type Stat struct {
	Value  int
	Total  int
	Method string
}

func NewStat(method string) *Stat {
	return &Stat{Method: method}
}

func (s *Stat) Analyze(data string) {
	res := strings.Contains(data, s.Method)

	s.Total++
	if res {
		s.Value++
	}
}

func (s *Stat) String() string {
	return fmt.Sprintf("Contains %d for string %s total lines %d\n", s.Value, s.Method, s.Total)
}
