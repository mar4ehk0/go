package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mar4ehk0/go/hw02_fix_app/typelist"
)

func ReadJSON(filePath string) ([]typelist.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []typelist.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}
