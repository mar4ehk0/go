package debug

import "fmt"

func Debug(value any) {
	fmt.Println(fmt.Sprintf("%#v", value))
}
