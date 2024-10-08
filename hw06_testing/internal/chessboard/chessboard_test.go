package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanNewChessboard(t *testing.T) {
	tests := []struct {
		name     string
		height   int
		width    int
		expected Board
	}{
		{
			"Chessboard: 0x0",
			0,
			0,
			Board{},
		},
		{
			"Chessboard: 0x1",
			0,
			1,
			Board{},
		},
		{
			"Chessboard: 1x1",
			1,
			1,
			Board{[]byte{35}},
		},
		{
			"Chessboard: 2x1",
			2,
			1,
			Board{[]byte{35}, []byte{32}},
		},
		{
			"Chessboard: 3x1",
			3,
			1,
			Board{[]byte{35}, []byte{32}, []byte{35}},
		},
		{
			"Chessboard: 3x3",
			3,
			3,
			Board{[]byte{35, 32, 35}, []byte{32, 35, 32}, []byte{35, 32, 35}},
		},
		{
			"Chessboard: 4x4",
			4,
			4,
			Board{
				[]byte{35, 32, 35, 32},
				[]byte{32, 35, 32, 35},
				[]byte{35, 32, 35, 32},
				[]byte{32, 35, 32, 35},
			},
		},
		{
			"Chessboard: 5x5",
			5,
			5,
			Board{
				[]byte{35, 32, 35, 32, 35},
				[]byte{32, 35, 32, 35, 32},
				[]byte{35, 32, 35, 32, 35},
				[]byte{32, 35, 32, 35, 32},
				[]byte{35, 32, 35, 32, 35},
			},
		},
		{
			"Chessboard: 6x6",
			6,
			6,
			Board{
				[]byte{35, 32, 35, 32, 35, 32},
				[]byte{32, 35, 32, 35, 32, 35},
				[]byte{35, 32, 35, 32, 35, 32},
				[]byte{32, 35, 32, 35, 32, 35},
				[]byte{35, 32, 35, 32, 35, 32},
				[]byte{32, 35, 32, 35, 32, 35},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := NewChessboard(tc.height, tc.width)
			assert.Equal(t, tc.expected, actual.Board())
		})
	}
}

func TestFailNewChessboard(t *testing.T) {
	tests := []struct {
		name     string
		height   int
		width    int
		expected error
	}{
		{
			"Chessboard: -1x-1",
			-1,
			-1,
			ErrHeightLessZero,
		},
		{
			"Chessboard: -1x10",
			-1,
			10,
			ErrHeightLessZero,
		},
		{
			"Chessboard: 1x-10",
			1,
			-10,
			ErrWidthLessZero,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewChessboard(tc.height, tc.width)
			assert.ErrorIs(t, tc.expected, err)
		})
	}
}
