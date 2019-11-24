package main

import (
	"fmt"
)

type SudokuBoard struct {
	board [9][9]int
}

func (s *SudokuBoard) populate(initial int) {
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			if r == 1 && c == 1 {
				s.board[r-1][c-1] = initial
			} else {
				for i := 1; i <= 9; i++ {
					if s.vHas(c, i) {
						continue
					} else {
						s.board[r-1][c-1] = i
						break
					}
				}
			}
		}
	}
}

func (s SudokuBoard) isComplete() bool {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if !s.vHas(i, j) || !s.hHas(i, j) {
				return false
			}
		}
	}
	return true
}

func (s SudokuBoard) printBoard() {
	fmt.Println("   1 2 3 4 5 6 7 8 9")
	for i, val := range s.board {
		fmt.Println(i+1, val)
	}
}

func (s *SudokuBoard) vHas(v, has int) bool {
	if v > 9 || v < 1 {
		panic(v)
	}
	for _, val := range s.board {
		res := val[v-1]
		if res == has {
			return true
		}

	}
	return false
}

func (s *SudokuBoard) hHas(h, has int) bool {
	if h > 9 || h < 1 {
		panic(h)
	}
	ints := s.board[h-1]
	for i := range ints {
		if i == has {
			return true
		}
	}
	return false
}

func (s SudokuBoard) sqHas(sq, has int) bool {
	panic("implement me")
}

type Board interface {
	vHas(v, has int) bool
	hHas(h, has int) bool
	sqHas(sq, has int) bool
	printBoard()
	isComplete() bool
	populate(initial int)
}

func main() {
	s := SudokuBoard{board: [9][9]int{{1}}}
	s.populate(3)
	s.printBoard()
	s.vHas(1, 1)
	fmt.Print(s.isComplete())
}
