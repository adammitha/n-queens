package main

import "fmt"

// nextBoards takes an existing board and creates a slice of new boards
// with each empty space filled with a queen.
// E.g.
// F F    T F  F T  F F  F F
// F F -> F F, F F, T F, F T
func (b *Board) nextBoards() []*Board {
	n := len(*b)

	boards := make([]*Board, 0)

	// Loop through b, fill every empty slot with a queen
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Check if slot already has a queen
			if (*b)[i][j] {
				continue
			}
			// Create a new board
			newBoard := b.duplicate()
			// Fill slot with a new queen
			newBoard[i][j] = true
			// Add board to list of boards if it is valid
			if newBoard.valid() {
				boards = append(boards, &newBoard)
			}
		}
	}

	return boards
}

// Solve consumes an n x n chess board and produces a pointer to
// a Board with n queens that cannot attack each other
func Solve(board *Board) (*Board, error) {
	// Check to see if board is solved
	if board.solved() {
		return board, nil
	}

	// Generate child boards
	nextBoards := board.nextBoards()

	// Loop through next boards
	for _, board := range nextBoards {
		solution, err := Solve(board)
		if err != nil {
			continue
		}

		return solution, nil
	}

	return nil, fmt.Errorf("unable to find a solution")
}
