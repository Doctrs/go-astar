package astar_go

type Map struct {
	Width   int16
	Height  int16
	Blocked map[Coordinates]struct{}
}

func (m *Map) CheckBlocked(c Coordinates) bool {
	_, ok := m.Blocked[c]

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
}

func (a *AStar) FindPath() []Coordinates {
	if a.To.X < 0 || a.To.Y < 0 {
		return make([]Coordinates, 0)
	}
	if a.To.X >= a.Map.Width || a.To.Y >= a.Map.Height {
		return make([]Coordinates, 0)
	}
	if a.Map.CheckBlocked(a.To) {
		return make([]Coordinates, 0)
	}
	if a.Map.CheckBlocked(a.From) {
		return make([]Coordinates, 0)
	}

	a.vectorMap = make(map[Coordinates]int)

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
	neighbour := map[Coordinates]struct{}{
		a.To: {},
	}
	for {
		if len(neighbour) == 0 {
			break
		}

		tempNeighbour := make(map[Coordinates]struct{})
		for coord := range neighbour {
			for _, c := range a.GetNeighbour(coord, false) {
				a.vectorMap[c] = i
				tempNeighbour[c] = struct{}{}
				i++

				if c == a.From {
					return
				}
			}
		}
		neighbour = tempNeighbour
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
		cord := Coordinates{X: c.X - 1, Y: c.Y}
		_, ok := a.vectorMap[cord]
		if !a.Map.CheckBlocked(cord) && (diag || !ok) {
			neighbour = append(neighbour, cord)
		}
	}
	if c.Y > 0 {
		cord := Coordinates{X: c.X, Y: c.Y - 1}
		_, ok := a.vectorMap[cord]
		if !a.Map.CheckBlocked(cord) && (diag || !ok) {
			neighbour = append(neighbour, cord)
		}
	}
	if c.X < a.Map.Width-1 {
		cord := Coordinates{X: c.X + 1, Y: c.Y}
		_, ok := a.vectorMap[cord]
		if !a.Map.CheckBlocked(cord) && (diag || !ok) {
			neighbour = append(neighbour, cord)
		}
	}
	if c.Y < a.Map.Height-1 {
		cord := Coordinates{X: c.X, Y: c.Y + 1}
		_, ok := a.vectorMap[cord]
		if !a.Map.CheckBlocked(cord) && (diag || !ok) {
			neighbour = append(neighbour, cord)
		}
	}

	if diag {
		if c.X > 0 && c.Y > 0 {
			cord := Coordinates{X: c.X - 1, Y: c.Y - 1}
			if !a.Map.CheckBlocked(cord) {
				neighbour = append(neighbour, cord)
			}
		}

		if c.X < a.Map.Width-1 && c.Y < a.Map.Height-1 {
			cord := Coordinates{X: c.X + 1, Y: c.Y + 1}
			if !a.Map.CheckBlocked(cord) {
				neighbour = append(neighbour, cord)
			}
		}

		if c.X > 0 && c.Y < a.Map.Height-1 {
			cord := Coordinates{X: c.X - 1, Y: c.Y + 1}
			if !a.Map.CheckBlocked(cord) {
				neighbour = append(neighbour, cord)
			}
		}

		if c.X < a.Map.Width-1 && c.Y > 0 {
			cord := Coordinates{X: c.X + 1, Y: c.Y - 1}
			if !a.Map.CheckBlocked(cord) {
				neighbour = append(neighbour, cord)
			}
		}
	}

	return neighbour
}
