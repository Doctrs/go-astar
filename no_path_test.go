package astar_go

import (
	"fmt"
	"reflect"
	"testing"
)

type TestNo struct {
	From   Coordinates
	To     Coordinates
	Fill   map[Coordinates]struct{}
}

func TestNoPath(t *testing.T) {
	cases := map[string]TestNo{}

	barier := make(map[Coordinates]struct{})
	for b := int16(0); b < 100; b++ {
		barier[Coordinates{
			X: 50,
			Y: b,
		}] = struct{}{}
	}
	case1 := TestNo{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 99},
		Fill:   barier,
	}

	cases["case1"] = case1

	barier = make(map[Coordinates]struct{})
	for b := int16(0); b < 100; b++ {
		barier[Coordinates{
			X: 50,
			Y: b,
		}] = struct{}{}
	}
	case2 := TestNo{
		From:   Coordinates{99, 99},
		To:     Coordinates{0, 0},
		Fill:   barier,
	}

	cases["case2"] = case2

	barier = make(map[Coordinates]struct{})
	for b := int16(0); b < 100; b++ {
		barier[Coordinates{
			X: 50,
			Y: b,
		}] = struct{}{}
	}
	case3 := TestNo{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 0},
		Fill:   barier,
	}
	cases["case3"] = case3

	barier = make(map[Coordinates]struct{})
	for b := int16(0); b < 100; b++ {
		barier[Coordinates{
			X: b,
			Y: 50,
		}] = struct{}{}
	}
	case4 := TestNo{
		From:   Coordinates{0, 0},
		To:     Coordinates{0, 99},
		Fill:   barier,
	}

	cases["case4"] = case4

	barier = make(map[Coordinates]struct{})
	for b := int16(99); b >= 0; b-- {
		barier[Coordinates{
			X: b,
			Y: 20,
		}] = struct{}{}
		barier[Coordinates{
			X: b,
			Y: 70,
		}] = struct{}{}
	}
	case5 := TestNo{
		From:   Coordinates{99, 99},
		To:     Coordinates{99, 0},
		Fill:   barier,
	}
	cases["case5"] = case5

	barier = make(map[Coordinates]struct{})
	for b := int16(99); b >= 0; b-- {
		barier[Coordinates{
			X: 20,
			Y: b,
		}] = struct{}{}
		barier[Coordinates{
			X: 70,
			Y: b,
		}] = struct{}{}
	}
	case6 := TestNo{
		From:   Coordinates{99, 99},
		To:     Coordinates{0, 99},
		Fill:   barier,
	}
	cases["case6"] = case6

	cases["case7"] = TestNo{
		From:   Coordinates{100, 0},
		To:     Coordinates{0, 0},
	}

	cases["case8"] = TestNo{
		From:   Coordinates{0, 100},
		To:     Coordinates{0, 0},
	}

	cases["case9"] = TestNo{
		From:   Coordinates{99, 99},
		To:     Coordinates{-1, 0},
	}

	cases["case10"] = TestNo{
		From:   Coordinates{99, 99},
		To:     Coordinates{-1, 0},
	}

	cases["case11"] = TestNo{
		From:   Coordinates{50, 50},
		To:     Coordinates{0, 0},
		Fill: map[Coordinates]struct{}{
			Coordinates{
				X: 50,
				Y: 50,
			}: {},
		},
	}
	cases["case12"] = TestNo{
		From:   Coordinates{0, 0},
		To:     Coordinates{50, 50},
		Fill: map[Coordinates]struct{}{
			Coordinates{
				X: 50,
				Y: 50,
			}: {},
		},
	}

	for i, testCase := range cases {
		astar := AStar{
			Map: Map{
				Width:   100,
				Height:  100,
				Blocked: testCase.Fill,
			},
			From: testCase.From,
			To:   testCase.To,
		}

		result := reflect.DeepEqual(make([]Coordinates, 0), astar.FindPath())
		if !result {
			t.Errorf("slice not equal in %s", i)
			fmt.Println(astar.FindPath())
		}
	}
}
