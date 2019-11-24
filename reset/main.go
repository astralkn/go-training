package main

import "fmt"

func main() {
	s := &BruteSudokuBoard{}
	s.populate(3)
	elm := s.Decisions.First()
	for {
		fmt.Printf("r:%d, c:%d, sq:%d, v:%d\n", elm.row, elm.column, elm.square, elm.value)
		elm = elm.Next()
		if elm == nil {
			break
		}
	}

}
