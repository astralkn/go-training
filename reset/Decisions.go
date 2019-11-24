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
	prev   *Decision
	next   *Decision
}

func (d Decisions) First() *Decision {
	return d.first
}

func (d Decisions) Last() *Decision {
	return d.first
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
		decision.prev = d.tail.next
		d.tail.next = decision
		d.tail = d.tail.next
	}
}
