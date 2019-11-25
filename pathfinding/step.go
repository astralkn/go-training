package pathfinding

type Step struct {
	north, east, south, west *Step
	prev                     *Step
	pos                      position
}

type position struct {
	x, y int
}

func (s Step) Init(prev *Step, dir string) *Step {
	step := new(Step)
	step.prev = prev
	switch dir {
	case "N":
		step.north = prev
		return step
	case "E":
		step.east = prev
		return step
	case "S":
		step.south = prev
		return step
	case "W":
		step.west = prev
		return step
	default:
		return step
	}
}

func (s *Step) GoNorth() *Step {
	return s.Init(s, "S")
}

func (s *Step) GoSouth() *Step {
	return s.Init(s, "N")
}
func (s *Step) GoEast() *Step {
	return s.Init(s, "W")
}
func (s *Step) GoNWest() *Step {
	return s.Init(s, "E")
}
