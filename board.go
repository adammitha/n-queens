package main

import "fmt"

// Board is a 2-dimensional array of booleans which represents a chess board
type Board [][]bool

// NewBoard returns a pointer to an empty n x n chess board
func NewBoard(n int) *Board {
	board := make(Board, n)

	for i := range board {
		board[i] = make([]bool, n)
	}

	return &board
}

// solved returns true if board contains n queens
// NB: assumes that board is valid (i.e. queens do not attack each other)
func (b *Board) solved() bool {
	n := len(*b)
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (*b)[i][j] {
				count++
			}
		}
	}
	return count == n
}

// valid returns true if all queens in the board cannot attack any other queen
func (b *Board) valid() bool {
	n := len(*b)
	// Find indices of all queens in the board
	queens := make([]Queen, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (*b)[i][j] {
				queens = append(queens, Queen{i, j})
			}
		}
	}

	for i, queen := range queens {
		if len(queens[i:]) == 0 {
			break
		}

		if queen.attacks(queens[i+1:]) {
			return false
		}
	}

	return true
}

// duplicate copies old board into a new board and returns the new board
func (b Board) duplicate() Board {
	duplicate := make(Board, len(b))

	for i := range b {
		duplicate[i] = make([]bool, len(b[i]))
		copy(duplicate[i], b[i])
	}

	return duplicate
}

// Render displays a picture of the chess board
func (b *Board) Render() {
	for _, row := range *b {
		var rowStr string
		for _, square := range row {
			if square {
				rowStr += " Q "
			} else {
				rowStr += " \u25a2 "
			}
		}
		fmt.Println(rowStr)
	}
}
