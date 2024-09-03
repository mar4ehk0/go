package view

import "fmt"

func PrintChessboard(chessboard []string) {
	for _, row := range chessboard {
		fmt.Println("|" + row + "|")
	}
}
