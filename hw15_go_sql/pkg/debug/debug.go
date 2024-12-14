package debug

import "fmt"

func Debug(value any) {
	fmt.Printf("DEBUG: %#v\n", value)
}
