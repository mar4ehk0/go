package hw02

import (
	"encoding/json"
	"io"
	"os"

	"github.com/mar4ehk0/go/hw06_testing/internal/hw02/employee"
)

const (
	pathToDefaultFile = "data.json"
)

func ReadFile(path string) ([]byte, error) {
	if len(path) == 0 {
		path = pathToDefaultFile
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	srcData, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return srcData, nil
}

func UnmarshalToEmployees(srcData []byte) ([]employee.Employee, error) {
	var employees []employee.Employee

	err := json.Unmarshal(srcData, &employees)

	return employees, err
}
