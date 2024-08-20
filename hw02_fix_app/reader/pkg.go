package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

import "github.com/fixme_my_friend/hw02_fix_app/type_list"

func ReadJSON(filePath string, limit int) ([]type_list.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []type_list.Employee

	err = json.Unmarshal(bytes, &data)

	res := data

	return res, nil
}
