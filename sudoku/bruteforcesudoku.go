package main

import (
	"errors"
	"fmt"
	"math"
)

type BruteSudokuBoard struct {
	board [9][9]int
	Decisions
}

func (s BruteSudokuBoard) getSquare(c, r int) int {
	row := math.Floor(float64(r) / 3)
	col := math.Ceil(float64(c+1) / 3)
	sq := int(row*3 + col)
	return sq
}

func (s BruteSudokuBoard) get(h, v int) int {
	return s.board[h][v]
}

func (s *BruteSudokuBoard) SolveGivenBoard(givenBoard [9][9]int) error {
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			givenValue := givenBoard[row][column]
			if givenValue > 0 && givenValue < 10 {
				d := &Decision{row, column, s.getSquare(column, row), givenValue, 0, nil, nil, true}
				s.Decisions.Push(d)
			} else if givenValue < 0 || givenValue > 9 {
				return errors.New("given values exceed limits")
			} else {
				d := &Decision{row, column, s.getSquare(column, row), 0, 0, nil, nil, false}
				s.Decisions.Push(d)
			}
		}
	}
	s.updateBoard()
	s.SolveBoard()
	return nil
}

func (s *BruteSudokuBoard) updateBoard() {
	n := s.Decisions.First()
	for {
		s.board[n.row][n.column] = n.value
		n = n.Next()
		if n == nil {
			break
		}
	}
}

func (s *BruteSudokuBoard) SolveBoard() {
	first := s.Decisions.First()
	for {
		if first.locked {
			first = first.Next()
		} else {
			break
		}
	}

	first.Populate(s)
}

func (s *BruteSudokuBoard) InitBoard() {
	for row := 0; row <= 8; row++ {
		for column := 0; column <= 8; column++ {
			d := &Decision{row, column, s.getSquare(column, row), 0, 0, nil, nil, false}
			s.Decisions.Push(d)
		}
	}
}

func (s BruteSudokuBoard) isComplete() bool {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if !s.vHas(i, j) || !s.hHas(i, j) {
				return false
			}
		}
	}
	return true
}

func (s *BruteSudokuBoard) PrintBoard() {

	fmt.Println("   1 2 3 4 5 6 7 8 9")
	for i, val := range s.board {
		fmt.Println(i+1, val)
	}
}

func (s BruteSudokuBoard) vHas(v, has int) bool {
	if v > 8 || v < 0 {
		panic(v)
	}
	result := false
	for i := 0; i < 9; i++ {
		if s.get(i, v) == has {
			result = true
		}
	}
	return result
}

func (s BruteSudokuBoard) hHas(h, has int) bool {
	if h > 8 || h < 0 {
		panic(h)
	}
	result := false
	for i := 0; i < 9; i++ {
		if s.get(h, i) == has {
			result = true
		}
	}
	return result
}

func (s BruteSudokuBoard) sqHas(sq, has int) bool {
	switch sq {
	case 1:
		return s.board[0][0] == has ||
			s.board[0][1] == has ||
			s.board[0][2] == has ||
			s.board[1][0] == has ||
			s.board[1][1] == has ||
			s.board[1][2] == has ||
			s.board[2][0] == has ||
			s.board[2][1] == has ||
			s.board[2][2] == has
	case 2:
		return s.board[0][3] == has ||
			s.board[0][4] == has ||
			s.board[0][5] == has ||
			s.board[1][3] == has ||
			s.board[1][4] == has ||
			s.board[1][5] == has ||
			s.board[2][3] == has ||
			s.board[2][4] == has ||
			s.board[2][5] == has
	case 3:
		return s.board[0][6] == has ||
			s.board[0][7] == has ||
			s.board[0][8] == has ||
			s.board[1][6] == has ||
			s.board[1][7] == has ||
			s.board[1][8] == has ||
			s.board[2][6] == has ||
			s.board[2][7] == has ||
			s.board[2][8] == has
	case 4:
		return s.board[3][0] == has ||
			s.board[3][1] == has ||
			s.board[3][2] == has ||
			s.board[4][0] == has ||
			s.board[4][1] == has ||
			s.board[4][2] == has ||
			s.board[5][0] == has ||
			s.board[5][1] == has ||
			s.board[5][2] == has
	case 5:
		return s.board[3][3] == has ||
			s.board[3][4] == has ||
			s.board[3][5] == has ||
			s.board[4][3] == has ||
			s.board[4][4] == has ||
			s.board[4][5] == has ||
			s.board[5][3] == has ||
			s.board[5][4] == has ||
			s.board[5][5] == has
	case 6:
		return s.board[3][6] == has ||
			s.board[3][7] == has ||
			s.board[3][8] == has ||
			s.board[4][6] == has ||
			s.board[4][7] == has ||
			s.board[4][8] == has ||
			s.board[5][6] == has ||
			s.board[5][7] == has ||
			s.board[5][8] == has
	case 7:
		return s.board[6][0] == has ||
			s.board[6][1] == has ||
			s.board[6][2] == has ||
			s.board[7][0] == has ||
			s.board[7][1] == has ||
			s.board[7][2] == has ||
			s.board[8][0] == has ||
			s.board[8][1] == has ||
			s.board[8][2] == has
	case 8:
		return s.board[6][3] == has ||
			s.board[6][4] == has ||
			s.board[6][5] == has ||
			s.board[7][3] == has ||
			s.board[7][4] == has ||
			s.board[7][5] == has ||
			s.board[8][3] == has ||
			s.board[8][4] == has ||
			s.board[8][5] == has
	case 9:
		return s.board[6][6] == has ||
			s.board[6][7] == has ||
			s.board[6][8] == has ||
			s.board[7][6] == has ||
			s.board[7][7] == has ||
			s.board[7][8] == has ||
			s.board[8][6] == has ||
			s.board[8][7] == has ||
			s.board[8][8] == has

	}
	panic(sq)
}
