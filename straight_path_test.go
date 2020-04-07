package astar_go

import (
	"fmt"
	"reflect"
	"testing"
)

type TestStraight struct {
	From   Coordinates
	To     Coordinates
	Result []Coordinates
}

func TestStraightPath(t *testing.T) {
	cases := map[string]TestStraight{}
	case1 := TestStraight{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 99},
		Result: nil,
	}
	res := make([]Coordinates, 0, 99)
	for i := int16(1); i < 100; i++ {
		res = append(res, Coordinates{
			X: i,
			Y: i,
		})
	}
	case1.Result = res
	cases["case1"] = case1

	case2 := TestStraight{
		From:   Coordinates{99, 99},
		To:     Coordinates{0, 0},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(98); i >= 0; i-- {
		res = append(res, Coordinates{
			X: i,
			Y: i,
		})
	}
	case2.Result = res
	cases["case2"] = case2

	case3 := TestStraight{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 0},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(1); i < 100; i++ {
		res = append(res, Coordinates{
			X: i,
			Y: 0,
		})
	}
	case3.Result = res
	cases["case3"] = case3

	case4 := TestStraight{
		From:   Coordinates{0, 0},
		To:     Coordinates{0, 99},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(1); i < 100; i++ {
		res = append(res, Coordinates{
			X: 0,
			Y: i,
		})
	}
	case4.Result = res
	cases["case4"] = case4

	case5 := TestStraight{
		From:   Coordinates{99, 99},
		To:     Coordinates{99, 0},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(98); i >= 0; i-- {
		res = append(res, Coordinates{
			X: 99,
			Y: i,
		})
	}
	case5.Result = res
	cases["case5"] = case5

	case6 := TestStraight{
		From:   Coordinates{99, 99},
		To:     Coordinates{0, 99},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(98); i >= 0; i-- {
		res = append(res, Coordinates{
			X: i,
			Y: 99,
		})
	}
	case6.Result = res
	cases["case6"] = case6

	case7 := TestStraight{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 50},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(1); i < 100; i++ {
		y := i
		if y > 50 {
			y = 50
		}
		res = append(res, Coordinates{
			X: i,
			Y: y,
		})
	}
	case7.Result = res
	cases["case7"] = case7

	case8 := TestStraight{
		From:   Coordinates{0, 0},
		To:     Coordinates{50, 99},
		Result: nil,
	}
	res = make([]Coordinates, 0, 99)
	for i := int16(1); i < 100; i++ {
		x := i
		if x > 50 {
			x = 50
		}
		res = append(res, Coordinates{
			X: x,
			Y: i,
		})
	}
	case8.Result = res
	cases["case8"] = case8



	for i, testCase := range cases {
		astar := AStar{
			Map: Map{
				Width:  100,
				Height: 100,
			},
			From: testCase.From,
			To: testCase.To,
		}

		result := reflect.DeepEqual(testCase.Result, astar.FindPath())
		if !result {
			t.Errorf("slice not equal in %s", i)
			fmt.Println(testCase.Result)
			fmt.Println(astar.FindPath())
		}
	}
}
