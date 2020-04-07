package astar_go

import "testing"

func BenchmarkLargeMap(b *testing.B) {
	astar := AStar{
		Map: Map{
			Width:  1000,
			Height: 1000,
		},
		From: Coordinates{
			X: 0,
			Y: 0,
		},
		To: Coordinates{
			X: 999,
			Y: 999,
		},
	}

	for i := 0; i < b.N; i++ {
		astar.FindPath()
	}
}
