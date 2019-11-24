package main

type SudokuBoard interface {
	vHas(v, has int) bool
	hHas(h, has int) bool
	sqHas(sq, has int) bool
	PrintBoard()
	isComplete() bool
	PopulateBoard(initial int)
	updateBoard()
}
