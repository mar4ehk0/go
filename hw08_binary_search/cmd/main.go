package main

import (
	"fmt"

	"github.com/mar4ehk0/hw08_binary_search/pkg/search"
)

func main() {
	values := []int{1, 2, 2, 2, 5, 6, 7}

	result, _ := search.BinarySearch(values, 2)

	fmt.Println(result)
}
