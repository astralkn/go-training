package main

type Decisions struct {
	first *Decision
	tail  *Decision
}

type Decision struct {
	row    int
	column int
	square int
	value  int
	strVal int
	prev   *Decision
	next   *Decision
	locked bool
}

func (d Decisions) First() *Decision {
	return d.first
}

func (d Decisions) Last() *Decision {
	return d.tail
}

func (d *Decision) Populate(b SudokuBoard) {
	d.strVal++
	done := false
	for ; d.strVal <= 9; d.strVal++ {
		if b.hHas(d.row, d.strVal) || b.vHas(d.column, d.strVal) || b.sqHas(d.square, d.strVal) {
			continue
		} else {
			d.value = d.strVal
			done = true
			break
		}
	}
	if done {
		b.updateBoard()
		nd := d.Next()
		for {
			if nd == nil {
				return
			}
			if nd.locked {
				nd = nd.Next()
			} else {
				break
			}
		}
		nd.Populate(b)
	} else {
		d.value = 0
		d.strVal = 0
		b.updateBoard()
		nd := d.Prev()
		for {
			if nd == nil {
				return
			}
			if nd.locked {
				nd = nd.Prev()
			} else {
				break
			}
		}
		nd.Populate(b)
	}
}

func (d Decision) Next() *Decision {
	return d.next
}

func (d Decision) Prev() *Decision {
	return d.prev
}

func (d *Decisions) Push(decision *Decision) {
	if d.first == nil {
		d.first = decision
		d.tail = d.first
	} else {
		decision.prev = d.tail
		d.tail.next = decision
		d.tail = d.tail.next
	}
}
