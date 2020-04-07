package astar_go

type Map struct {
	Width   int16
	Height  int16
	Blocked map[Coordinates]struct{}
}

func (m *Map) CheckFill(x int16, y int16) bool {
	_, ok := m.Blocked[Coordinates{x, y}]

	return ok
}

type Coordinates struct {
	X int16
	Y int16
}

type AStar struct {
	Map       Map
	From      Coordinates
	To        Coordinates
	vectorMap map[Coordinates]int
	neighbour map[Coordinates]struct{}
}

func (a *AStar) FindPath() []Coordinates {
	if a.To.X < 0 || a.To.Y < 0 {
		return make([]Coordinates, 0)
	}
	if a.To.X >= a.Map.Width || a.To.Y >= a.Map.Height {
		return make([]Coordinates, 0)
	}
	if a.Map.CheckFill(a.To.X, a.To.Y) {
		return make([]Coordinates, 0)
	}
	if a.Map.CheckFill(a.From.X, a.From.Y) {
		return make([]Coordinates, 0)
	}

	a.neighbour = make(map[Coordinates]struct{})
	a.vectorMap = make(map[Coordinates]int)

	a.neighbour[a.To] = struct{}{}
	a.LoopFind()
	a.vectorMap[a.To] = 0

	_, ok := a.vectorMap[a.From]
	if !ok {
		return make([]Coordinates, 0)
	}

	path := make([]Coordinates, 0)

	coords := a.From
	for {
		coords = a.GetMinNeighbour(coords)
		path = append(path, coords)

		if coords == a.To {
			break
		}
	}

	return path
}

func (a *AStar) LoopFind() {
	i := 1
	for {
		if len(a.neighbour) == 0 {
			break
		}

		tempNeighbour := make(map[Coordinates]struct{})
		for coord := range a.neighbour {
			for _, c := range a.GetNeighbour(coord, false) {
				_, ok := a.vectorMap[c]
				if ok {
					continue
				}

				a.vectorMap[c] = i
				tempNeighbour[c] = struct{}{}
				i++

				if c == a.From {
					return
				}
			}
		}
		a.neighbour = tempNeighbour
	}
}

func (a *AStar) GetMinNeighbour(c Coordinates) (result Coordinates) {
	var m int
	i := true
	for _, e := range a.GetNeighbour(c, true) {
		_, ok := a.vectorMap[e]
		if !ok {
			continue
		}

		if i == true || a.vectorMap[e] < m {
			i = false
			m = a.vectorMap[e]
			result = e
		}
	}

	return result
}

func (a *AStar) GetNeighbour(c Coordinates, diag bool) (neighbour []Coordinates) {
	if c.X > 0 {
		if !a.Map.CheckFill(c.X-1, c.Y) {
			neighbour = append(neighbour, Coordinates{X: c.X - 1, Y: c.Y})
		}
	}
	if c.Y > 0 {
		if !a.Map.CheckFill(c.X, c.Y-1) {
			neighbour = append(neighbour, Coordinates{X: c.X, Y: c.Y - 1})
		}
	}
	if c.X < a.Map.Width-1 {
		if !a.Map.CheckFill(c.X+1, c.Y) {
			neighbour = append(neighbour, Coordinates{X: c.X + 1, Y: c.Y})
		}
	}
	if c.Y < a.Map.Height-1 {
		if !a.Map.CheckFill(c.X, c.Y+1) {
			neighbour = append(neighbour, Coordinates{X: c.X, Y: c.Y + 1})
		}
	}

	if diag {
		if c.X > 0 && c.Y > 0 {
			if !a.Map.CheckFill(c.X-1, c.Y-1) {
				neighbour = append(neighbour, Coordinates{X: c.X - 1, Y: c.Y - 1})
			}
		}

		if c.X < a.Map.Width-1 && c.Y < a.Map.Height-1 {
			if !a.Map.CheckFill(c.X+1, c.Y+1) {
				neighbour = append(neighbour, Coordinates{X: c.X + 1, Y: c.Y + 1})
			}
		}

		if c.X > 0 && c.Y < a.Map.Height-1 {
			if !a.Map.CheckFill(c.X-1, c.Y+1) {
				neighbour = append(neighbour, Coordinates{X: c.X - 1, Y: c.Y + 1})
			}
		}

		if c.X < a.Map.Width-1 && c.Y > 0 {
			if !a.Map.CheckFill(c.X+1, c.Y-1) {
				neighbour = append(neighbour, Coordinates{X: c.X + 1, Y: c.Y - 1})
			}
		}
	}

	return neighbour
}
