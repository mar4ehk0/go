package chessboard

import (
	"errors"
	"strings"
)

type Chessboard struct {
	height int
	width  int
	board  Board
}

type (
	rowChessboard []byte
	Board         []rowChessboard
)

const (
	BlackCeil = 35
	WhiteCeil = 32
)

var (
	ErrHeightLessZero = errors.New("height less zero")
	ErrWidthLessZero  = errors.New("width less zero")
)

func NewChessboard(height int, width int) (*Chessboard, error) {
	err := validate(height, width)
	if err != nil {
		return nil, err
	}

	board := make([]rowChessboard, height)
	chessboard := &Chessboard{height: height, width: width, board: board}
	chessboard.fillChessboard()

	return chessboard, nil
}

func validate(height, width int) error {
	if height < 0 {
		return ErrHeightLessZero
	}
	if width < 0 {
		return ErrWidthLessZero
	}
	return nil
}

func (c *Chessboard) fillChessboard() {
	rowEven := createEvenRow(c.width)
	rowOdd := createOddRow(rowEven)

	for i := 0; i < len(c.board); i++ {
		row := make(rowChessboard, c.width)
		if i%2 == 0 {
			copy(row, rowEven)
		} else {
			copy(row, rowOdd)
		}
		c.board[i] = row
	}
}

func (c *Chessboard) Board() Board {
	return c.board
}

func (c *Chessboard) String() string {
	rowBuilder := &strings.Builder{}
	for i := 0; i < len(c.board); i++ {
		rowBuilder.Write(c.board[i])
		rowBuilder.WriteRune('\n')
	}

	return rowBuilder.String()
}

func createEvenRow(width int) rowChessboard {
	row := make(rowChessboard, width)

	for i := 0; i < width; i++ {
		if i%2 == 0 {
			row[i] = BlackCeil
		} else {
			row[i] = WhiteCeil
		}
	}

	return row
}

func createOddRow(rowEven rowChessboard) rowChessboard {
	rowOdd := make(rowChessboard, len(rowEven))

	for i := 0; i < len(rowEven); i++ {
		if rowEven[i] == BlackCeil {
			rowOdd[i] = WhiteCeil
		} else {
			rowOdd[i] = BlackCeil
		}
	}

	return rowOdd
}
