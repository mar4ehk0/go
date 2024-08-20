package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/type_list"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []type_list.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
