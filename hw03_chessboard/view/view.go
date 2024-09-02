package view

import "fmt"

func PrintChessboard(chessboard [][]string) {
	for _, row := range chessboard {
		fmt.Print("|")
		for _, value := range row {
			fmt.Print(value)
		}
		fmt.Println("|")
	}
}
