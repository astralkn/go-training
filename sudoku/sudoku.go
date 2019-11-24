package main

type SudokuBoard interface {
	SolveGivenBoard(givenBoard [9][9]int) error
	vHas(v, has int) bool
	hHas(h, has int) bool
	sqHas(sq, has int) bool
	PrintBoard()
	isComplete() bool
	SolveBoard()
	updateBoard()
}
