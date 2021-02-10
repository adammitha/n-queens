package main

import "math"

// Queen represents position (row, column) of a Queen on a chess board
type Queen [2]int

// attacks checks if q attacks any of the queens
func (q Queen) attacks(queens []Queen) bool {
	for _, queen := range queens {
		if q[0] == queen[0] || q[1] == queen[1] || q.diagonal(queen) {
			return true
		}
	}
	return false
}

func (q Queen) diagonal(queen Queen) bool {
	dx := int(math.Abs(float64(q[1] - queen[1])))
	dy := int(math.Abs(float64(q[0] - queen[0])))

	return dx == dy
}
