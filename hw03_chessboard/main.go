package main

import (
	"fmt"

	"github.com/mar4ehk0/go/hw03_chessboard/validator"
	"github.com/mar4ehk0/go/hw03_chessboard/view"
)

const (
	BlackCeil               = "#"
	WhiteCeil               = " "
	ChessboardDefaultHeight = 8
	ChessboardDefaultWidth  = 8
)

func main() {
	var height int
	var width int

	fmt.Println("Please input chessboard size: ")
	fmt.Print("height: ")
	fmt.Scanln(&height)
	fmt.Print("width: ")
	fmt.Scanln(&width)

	if !validator.ValidateHeightChessboard(height) {
		height = ChessboardDefaultHeight
	}

	if !validator.ValidateWidthChessboard(width) {
		width = ChessboardDefaultWidth
	}

	chessboard := createChessboard(height, width)

	view.PrintChessboard(chessboard)
}

func createChessboard(height int, width int) [][]string {
	rowEven := createEvenRow(width)
	rowOdd := createOddRow(width)

	chessboard := make([][]string, height)

	for i := range chessboard {
		var row []string
		if i%2 == 0 {
			row = rowEven
		} else {
			row = rowOdd
		}
		chessboard[i] = row
	}

	return chessboard
}

func createEvenRow(width int) []string {
	return createRow(width, 1)
}

func createOddRow(width int) []string {
	return createRow(width, 0)
}

func createRow(width int, isEven int) []string {
	row := make([]string, width)
	for i := 0; i < width; i++ {
		var ceil string
		if i%2 == isEven {
			ceil = BlackCeil
		} else {
			ceil = WhiteCeil
		}

		row[i] = ceil
	}

	return row
}
