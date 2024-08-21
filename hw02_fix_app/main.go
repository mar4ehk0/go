package main

import (
	"fmt"

	"github.com/mar4ehk0/go/hw02_fix_app/printer"
	"github.com/mar4ehk0/go/hw02_fix_app/reader"
	"github.com/mar4ehk0/go/hw02_fix_app/typelist"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []typelist.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		printer.PrintStaff(staff)
	}
}
