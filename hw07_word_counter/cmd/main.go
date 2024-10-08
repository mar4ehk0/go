package main

import (
	"github.com/mar4ehk0/hw07_word_counter/internal/counter"
	"github.com/mar4ehk0/hw07_word_counter/pkg/debug"
)

func main() {
	source := "Hello John, Hello World. Hello world!world world:)world"
	words := counter.CountWords(source)

	debug.Debug(words)
}
