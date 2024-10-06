package hw02

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanReadFileWhenEmptyFilePath(t *testing.T) {
	expected := []byte(`
	[
    {
        "userId": 10,
        "age": 25,
        "name": "Rob",
        "departmentId": 3
    },
    {
        "userId": 11,
        "age": 30,
        "name": "George",
        "departmentId": 2
    }
]`)

	actual, _ := ReadFile("")

	require.JSONEq(t, string(expected), string(actual))
}

func TestCanUnmarshalToEmployees(t *testing.T) {
	tests := []struct {
		name     string
		srcData  []byte
		expected string
	}{
		{
			"1 item from JSON",
			[]byte(`[
				{
					"userId": 10,
					"age": 25,
					"name": "Rob",
					"departmentId": 3
				}
			]`),
			"User ID: 10; Age: 25; Name: Rob; Department ID: 3;",
		},
		{
			"Some items from JSON",
			[]byte(`[
				{
					"userId": 10,
					"age": 25,
					"name": "Rob",
					"departmentId": 3
				},
				{
					"userId": 11,
					"age": 26,
					"name": "John",
					"departmentId": 4
				}
			]`),
			"User ID: 10; Age: 25; Name: Rob; Department ID: 3;User ID: 11; Age: 26; Name: John; Department ID: 4;",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			rowBuilder := &strings.Builder{}

			employees, _ := UnmarshalToEmployees(tc.srcData)

			for _, v := range employees {
				str := fmt.Sprint(v)
				rowBuilder.WriteString(str)
			}

			assert.Equal(t, tc.expected, rowBuilder.String())
		})
	}
}

func TestFailUnmarshalToEmployeesWhenWrongJson(t *testing.T) {
	expected := &json.SyntaxError{}
	srcData := []byte(`[
				{
					"userId": 10,
					"age" 25,
					"name" "Rob",
					"departmentId": 3
				}
			]`)

	_, err := UnmarshalToEmployees(srcData)

	assert.IsType(t, expected, err)
}
