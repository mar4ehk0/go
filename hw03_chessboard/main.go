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
	var ceil string
	var prevCeil string

	chessboard := make([][]string, height)

	for i := range chessboard {
		chessboard[i] = make([]string, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			changeColorCeil(prevCeil, &ceil)

			chessboard[i][j] = ceil
			prevCeil = ceil
		}
		changeColorCeil(prevCeil, &prevCeil)
	}

	return chessboard
}

func changeColorCeil(value string, subject *string) {
	if value == BlackCeil {
		*subject = WhiteCeil
	} else {
		*subject = BlackCeil
	}
}
