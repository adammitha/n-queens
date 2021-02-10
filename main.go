package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatal("Please supply the size of the chess board as a single integer")
	}
	n, _ := strconv.Atoi(os.Args[1])

	emptyBoard := NewBoard(n)
	sol, err := Solve(emptyBoard)
	if err != nil {
		fmt.Println(err)
	} else {
		sol.Render()
	}

}
