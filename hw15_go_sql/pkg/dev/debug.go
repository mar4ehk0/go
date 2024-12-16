package dev

import "fmt"

func Debug(value any) {
	fmt.Printf("DEBUG: value - %#v, pointer - %p\n", value, value)
}
