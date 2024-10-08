package debug

import "fmt"

func Debug(value any) {
	fmt.Printf("%#v\n", value)
}
