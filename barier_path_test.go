package astar

import (
	"fmt"
	"reflect"
	"testing"
)

type TestBarrier struct {
	From   Coordinates
	To     Coordinates
	Result []Coordinates
	Fill   map[Coordinates]struct{}
}

func TestBarrierPath(t *testing.T) {
	cases := map[string]TestBarrier{}

	barier := make(map[Coordinates]struct{})
	for b := int16(0); b < 90; b++ {
		barier[Coordinates{
			X: 50,
			Y: b,
		}] = struct{}{}
	}
	case1 := TestBarrier{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 99},
		Fill:   barier,
		Result: nil,
	}

	coords := [139][2]int16{[2]int16{1, 1}, [2]int16{2, 2}, [2]int16{3, 3}, [2]int16{4, 4}, [2]int16{5, 5}, [2]int16{6, 6}, [2]int16{7, 7}, [2]int16{8, 8}, [2]int16{9, 9}, [2]int16{10, 10}, [2]int16{11, 11}, [2]int16{12, 12}, [2]int16{13, 13}, [2]int16{14, 14}, [2]int16{15, 15}, [2]int16{16, 16}, [2]int16{17, 17}, [2]int16{18, 18}, [2]int16{19, 19}, [2]int16{20, 20}, [2]int16{21, 21}, [2]int16{22, 22}, [2]int16{23, 23}, [2]int16{24, 24}, [2]int16{25, 25}, [2]int16{26, 26}, [2]int16{27, 27}, [2]int16{28, 28}, [2]int16{29, 29}, [2]int16{30, 30}, [2]int16{31, 31}, [2]int16{32, 32}, [2]int16{33, 33}, [2]int16{34, 34}, [2]int16{35, 35}, [2]int16{36, 36}, [2]int16{37, 37}, [2]int16{38, 38}, [2]int16{39, 39}, [2]int16{40, 40}, [2]int16{41, 41}, [2]int16{42, 42}, [2]int16{43, 43}, [2]int16{44, 44}, [2]int16{45, 45}, [2]int16{46, 46}, [2]int16{47, 47}, [2]int16{48, 48}, [2]int16{49, 49}, [2]int16{49, 50}, [2]int16{49, 51}, [2]int16{49, 52}, [2]int16{49, 53}, [2]int16{49, 54}, [2]int16{49, 55}, [2]int16{49, 56}, [2]int16{49, 57}, [2]int16{49, 58}, [2]int16{49, 59}, [2]int16{49, 60}, [2]int16{49, 61}, [2]int16{49, 62}, [2]int16{49, 63}, [2]int16{49, 64}, [2]int16{49, 65}, [2]int16{49, 66}, [2]int16{49, 67}, [2]int16{49, 68}, [2]int16{49, 69}, [2]int16{49, 70}, [2]int16{49, 71}, [2]int16{49, 72}, [2]int16{49, 73}, [2]int16{49, 74}, [2]int16{49, 75}, [2]int16{49, 76}, [2]int16{49, 77}, [2]int16{49, 78}, [2]int16{49, 79}, [2]int16{49, 80}, [2]int16{49, 81}, [2]int16{49, 82}, [2]int16{49, 83}, [2]int16{49, 84}, [2]int16{49, 85}, [2]int16{49, 86}, [2]int16{49, 87}, [2]int16{49, 88}, [2]int16{49, 89}, [2]int16{50, 90}, [2]int16{51, 91}, [2]int16{52, 92}, [2]int16{53, 93}, [2]int16{54, 94}, [2]int16{55, 95}, [2]int16{56, 96}, [2]int16{57, 97}, [2]int16{58, 98}, [2]int16{59, 99}, [2]int16{60, 99}, [2]int16{61, 99}, [2]int16{62, 99}, [2]int16{63, 99}, [2]int16{64, 99}, [2]int16{65, 99}, [2]int16{66, 99}, [2]int16{67, 99}, [2]int16{68, 99}, [2]int16{69, 99}, [2]int16{70, 99}, [2]int16{71, 99}, [2]int16{72, 99}, [2]int16{73, 99}, [2]int16{74, 99}, [2]int16{75, 99}, [2]int16{76, 99}, [2]int16{77, 99}, [2]int16{78, 99}, [2]int16{79, 99}, [2]int16{80, 99}, [2]int16{81, 99}, [2]int16{82, 99}, [2]int16{83, 99}, [2]int16{84, 99}, [2]int16{85, 99}, [2]int16{86, 99}, [2]int16{87, 99}, [2]int16{88, 99}, [2]int16{89, 99}, [2]int16{90, 99}, [2]int16{91, 99}, [2]int16{92, 99}, [2]int16{93, 99}, [2]int16{94, 99}, [2]int16{95, 99}, [2]int16{96, 99}, [2]int16{97, 99}, [2]int16{98, 99}, [2]int16{99, 99},}
	res := make([]Coordinates, 0, 99)
	for _, i := range coords {
		res = append(res, Coordinates{
			X: i[0],
			Y: i[1],
		})
	}
	case1.Result = res
	cases["case1"] = case1

	barier = make(map[Coordinates]struct{})
	for b := int16(0); b < 90; b++ {
		barier[Coordinates{
			X: 50,
			Y: b,
		}] = struct{}{}
	}
	case2 := TestBarrier{
		From:   Coordinates{99, 99},
		To:     Coordinates{0, 0},
		Result: nil,
		Fill:   barier,
	}

	coords = [139][2]int16{[2]int16{98, 98}, [2]int16{97, 97}, [2]int16{96, 96}, [2]int16{95, 95}, [2]int16{94, 94}, [2]int16{93, 93}, [2]int16{92, 92}, [2]int16{91, 91}, [2]int16{90, 90}, [2]int16{89, 90}, [2]int16{88, 90}, [2]int16{87, 90}, [2]int16{86, 90}, [2]int16{85, 90}, [2]int16{84, 90}, [2]int16{83, 90}, [2]int16{82, 90}, [2]int16{81, 90}, [2]int16{80, 90}, [2]int16{79, 90}, [2]int16{78, 90}, [2]int16{77, 90}, [2]int16{76, 90}, [2]int16{75, 90}, [2]int16{74, 90}, [2]int16{73, 90}, [2]int16{72, 90}, [2]int16{71, 90}, [2]int16{70, 90}, [2]int16{69, 90}, [2]int16{68, 90}, [2]int16{67, 90}, [2]int16{66, 90}, [2]int16{65, 90}, [2]int16{64, 90}, [2]int16{63, 90}, [2]int16{62, 90}, [2]int16{61, 90}, [2]int16{60, 90}, [2]int16{59, 90}, [2]int16{58, 90}, [2]int16{57, 90}, [2]int16{56, 90}, [2]int16{55, 90}, [2]int16{54, 90}, [2]int16{53, 90}, [2]int16{52, 90}, [2]int16{51, 90}, [2]int16{50, 90}, [2]int16{49, 89}, [2]int16{48, 88}, [2]int16{47, 87}, [2]int16{46, 86}, [2]int16{45, 85}, [2]int16{44, 84}, [2]int16{43, 83}, [2]int16{42, 82}, [2]int16{41, 81}, [2]int16{40, 80}, [2]int16{39, 79}, [2]int16{38, 78}, [2]int16{37, 77}, [2]int16{36, 76}, [2]int16{35, 75}, [2]int16{34, 74}, [2]int16{33, 73}, [2]int16{32, 72}, [2]int16{31, 71}, [2]int16{30, 70}, [2]int16{29, 69}, [2]int16{28, 68}, [2]int16{27, 67}, [2]int16{26, 66}, [2]int16{25, 65}, [2]int16{24, 64}, [2]int16{23, 63}, [2]int16{22, 62}, [2]int16{21, 61}, [2]int16{20, 60}, [2]int16{19, 59}, [2]int16{18, 58}, [2]int16{17, 57}, [2]int16{16, 56}, [2]int16{15, 55}, [2]int16{14, 54}, [2]int16{13, 53}, [2]int16{12, 52}, [2]int16{11, 51}, [2]int16{10, 50}, [2]int16{9, 49}, [2]int16{8, 48}, [2]int16{7, 47}, [2]int16{6, 46}, [2]int16{5, 45}, [2]int16{4, 44}, [2]int16{3, 43}, [2]int16{2, 42}, [2]int16{1, 41}, [2]int16{0, 40}, [2]int16{0, 39}, [2]int16{0, 38}, [2]int16{0, 37}, [2]int16{0, 36}, [2]int16{0, 35}, [2]int16{0, 34}, [2]int16{0, 33}, [2]int16{0, 32}, [2]int16{0, 31}, [2]int16{0, 30}, [2]int16{0, 29}, [2]int16{0, 28}, [2]int16{0, 27}, [2]int16{0, 26}, [2]int16{0, 25}, [2]int16{0, 24}, [2]int16{0, 23}, [2]int16{0, 22}, [2]int16{0, 21}, [2]int16{0, 20}, [2]int16{0, 19}, [2]int16{0, 18}, [2]int16{0, 17}, [2]int16{0, 16}, [2]int16{0, 15}, [2]int16{0, 14}, [2]int16{0, 13}, [2]int16{0, 12}, [2]int16{0, 11}, [2]int16{0, 10}, [2]int16{0, 9}, [2]int16{0, 8}, [2]int16{0, 7}, [2]int16{0, 6}, [2]int16{0, 5}, [2]int16{0, 4}, [2]int16{0, 3}, [2]int16{0, 2}, [2]int16{0, 1}, [2]int16{0, 0},}
	res = make([]Coordinates, 0, 99)
	for _, i := range coords {
		res = append(res, Coordinates{
			X: i[0],
			Y: i[1],
		})
	}
	case2.Result = res
	cases["case2"] = case2

	barier = make(map[Coordinates]struct{})
	for b := int16(0); b < 90; b++ {
		barier[Coordinates{
			X: 50,
			Y: b,
		}] = struct{}{}
	}
	case3 := TestBarrier{
		From:   Coordinates{0, 0},
		To:     Coordinates{99, 0},
		Result: nil,
		Fill:   barier,
	}
	coords1 := [180][2]int16{[2]int16{1, 1}, [2]int16{2, 2}, [2]int16{3, 3}, [2]int16{4, 4}, [2]int16{5, 5}, [2]int16{6, 6}, [2]int16{7, 7}, [2]int16{8, 8}, [2]int16{9, 9}, [2]int16{10, 10}, [2]int16{11, 11}, [2]int16{12, 12}, [2]int16{13, 13}, [2]int16{14, 14}, [2]int16{15, 15}, [2]int16{16, 16}, [2]int16{17, 17}, [2]int16{18, 18}, [2]int16{19, 19}, [2]int16{20, 20}, [2]int16{21, 21}, [2]int16{22, 22}, [2]int16{23, 23}, [2]int16{24, 24}, [2]int16{25, 25}, [2]int16{26, 26}, [2]int16{27, 27}, [2]int16{28, 28}, [2]int16{29, 29}, [2]int16{30, 30}, [2]int16{31, 31}, [2]int16{32, 32}, [2]int16{33, 33}, [2]int16{34, 34}, [2]int16{35, 35}, [2]int16{36, 36}, [2]int16{37, 37}, [2]int16{38, 38}, [2]int16{39, 39}, [2]int16{40, 40}, [2]int16{41, 41}, [2]int16{42, 42}, [2]int16{43, 43}, [2]int16{44, 44}, [2]int16{45, 45}, [2]int16{46, 46}, [2]int16{47, 47}, [2]int16{48, 48}, [2]int16{49, 49}, [2]int16{49, 50}, [2]int16{49, 51}, [2]int16{49, 52}, [2]int16{49, 53}, [2]int16{49, 54}, [2]int16{49, 55}, [2]int16{49, 56}, [2]int16{49, 57}, [2]int16{49, 58}, [2]int16{49, 59}, [2]int16{49, 60}, [2]int16{49, 61}, [2]int16{49, 62}, [2]int16{49, 63}, [2]int16{49, 64}, [2]int16{49, 65}, [2]int16{49, 66}, [2]int16{49, 67}, [2]int16{49, 68}, [2]int16{49, 69}, [2]int16{49, 70}, [2]int16{49, 71}, [2]int16{49, 72}, [2]int16{49, 73}, [2]int16{49, 74}, [2]int16{49, 75}, [2]int16{49, 76}, [2]int16{49, 77}, [2]int16{49, 78}, [2]int16{49, 79}, [2]int16{49, 80}, [2]int16{49, 81}, [2]int16{49, 82}, [2]int16{49, 83}, [2]int16{49, 84}, [2]int16{49, 85}, [2]int16{49, 86}, [2]int16{49, 87}, [2]int16{49, 88}, [2]int16{49, 89}, [2]int16{50, 90}, [2]int16{51, 89}, [2]int16{52, 88}, [2]int16{53, 87}, [2]int16{54, 86}, [2]int16{55, 85}, [2]int16{56, 84}, [2]int16{57, 83}, [2]int16{58, 82}, [2]int16{59, 81}, [2]int16{60, 80}, [2]int16{61, 79}, [2]int16{62, 78}, [2]int16{63, 77}, [2]int16{64, 76}, [2]int16{65, 75}, [2]int16{66, 74}, [2]int16{67, 73}, [2]int16{68, 72}, [2]int16{69, 71}, [2]int16{70, 70}, [2]int16{71, 69}, [2]int16{72, 68}, [2]int16{73, 67}, [2]int16{74, 66}, [2]int16{75, 65}, [2]int16{76, 64}, [2]int16{77, 63}, [2]int16{78, 62}, [2]int16{79, 61}, [2]int16{80, 60}, [2]int16{81, 59}, [2]int16{82, 58}, [2]int16{83, 57}, [2]int16{84, 56}, [2]int16{85, 55}, [2]int16{86, 54}, [2]int16{87, 53}, [2]int16{88, 52}, [2]int16{89, 51}, [2]int16{90, 50}, [2]int16{91, 49}, [2]int16{92, 48}, [2]int16{93, 47}, [2]int16{94, 46}, [2]int16{95, 45}, [2]int16{96, 44}, [2]int16{97, 43}, [2]int16{98, 42}, [2]int16{99, 41}, [2]int16{99, 40}, [2]int16{99, 39}, [2]int16{99, 38}, [2]int16{99, 37}, [2]int16{99, 36}, [2]int16{99, 35}, [2]int16{99, 34}, [2]int16{99, 33}, [2]int16{99, 32}, [2]int16{99, 31}, [2]int16{99, 30}, [2]int16{99, 29}, [2]int16{99, 28}, [2]int16{99, 27}, [2]int16{99, 26}, [2]int16{99, 25}, [2]int16{99, 24}, [2]int16{99, 23}, [2]int16{99, 22}, [2]int16{99, 21}, [2]int16{99, 20}, [2]int16{99, 19}, [2]int16{99, 18}, [2]int16{99, 17}, [2]int16{99, 16}, [2]int16{99, 15}, [2]int16{99, 14}, [2]int16{99, 13}, [2]int16{99, 12}, [2]int16{99, 11}, [2]int16{99, 10}, [2]int16{99, 9}, [2]int16{99, 8}, [2]int16{99, 7}, [2]int16{99, 6}, [2]int16{99, 5}, [2]int16{99, 4}, [2]int16{99, 3}, [2]int16{99, 2}, [2]int16{99, 1}, [2]int16{99, 0},}
	res = make([]Coordinates, 0, 99)
	for _, i := range coords1 {
		res = append(res, Coordinates{
			X: i[0],
			Y: i[1],
		})
	}
	case3.Result = res
	cases["case3"] = case3

	barier = make(map[Coordinates]struct{})
	for b := int16(0); b < 90; b++ {
		barier[Coordinates{
			X: b,
			Y: 50,
		}] = struct{}{}
	}
	case4 := TestBarrier{
		From:   Coordinates{0, 0},
		To:     Coordinates{0, 99},
		Result: nil,
		Fill:   barier,
	}
	coords1 = [180][2]int16{[2]int16{1, 1}, [2]int16{2, 2}, [2]int16{3, 3}, [2]int16{4, 4}, [2]int16{5, 5}, [2]int16{6, 6}, [2]int16{7, 7}, [2]int16{8, 8}, [2]int16{9, 9}, [2]int16{10, 10}, [2]int16{11, 11}, [2]int16{12, 12}, [2]int16{13, 13}, [2]int16{14, 14}, [2]int16{15, 15}, [2]int16{16, 16}, [2]int16{17, 17}, [2]int16{18, 18}, [2]int16{19, 19}, [2]int16{20, 20}, [2]int16{21, 21}, [2]int16{22, 22}, [2]int16{23, 23}, [2]int16{24, 24}, [2]int16{25, 25}, [2]int16{26, 26}, [2]int16{27, 27}, [2]int16{28, 28}, [2]int16{29, 29}, [2]int16{30, 30}, [2]int16{31, 31}, [2]int16{32, 32}, [2]int16{33, 33}, [2]int16{34, 34}, [2]int16{35, 35}, [2]int16{36, 36}, [2]int16{37, 37}, [2]int16{38, 38}, [2]int16{39, 39}, [2]int16{40, 40}, [2]int16{41, 41}, [2]int16{42, 42}, [2]int16{43, 43}, [2]int16{44, 44}, [2]int16{45, 45}, [2]int16{46, 46}, [2]int16{47, 47}, [2]int16{48, 48}, [2]int16{49, 49}, [2]int16{50, 49}, [2]int16{51, 49}, [2]int16{52, 49}, [2]int16{53, 49}, [2]int16{54, 49}, [2]int16{55, 49}, [2]int16{56, 49}, [2]int16{57, 49}, [2]int16{58, 49}, [2]int16{59, 49}, [2]int16{60, 49}, [2]int16{61, 49}, [2]int16{62, 49}, [2]int16{63, 49}, [2]int16{64, 49}, [2]int16{65, 49}, [2]int16{66, 49}, [2]int16{67, 49}, [2]int16{68, 49}, [2]int16{69, 49}, [2]int16{70, 49}, [2]int16{71, 49}, [2]int16{72, 49}, [2]int16{73, 49}, [2]int16{74, 49}, [2]int16{75, 49}, [2]int16{76, 49}, [2]int16{77, 49}, [2]int16{78, 49}, [2]int16{79, 49}, [2]int16{80, 49}, [2]int16{81, 49}, [2]int16{82, 49}, [2]int16{83, 49}, [2]int16{84, 49}, [2]int16{85, 49}, [2]int16{86, 49}, [2]int16{87, 49}, [2]int16{88, 49}, [2]int16{89, 49}, [2]int16{90, 50}, [2]int16{89, 51}, [2]int16{88, 52}, [2]int16{87, 53}, [2]int16{86, 54}, [2]int16{85, 55}, [2]int16{84, 56}, [2]int16{83, 57}, [2]int16{82, 58}, [2]int16{81, 59}, [2]int16{80, 60}, [2]int16{79, 61}, [2]int16{78, 62}, [2]int16{77, 63}, [2]int16{76, 64}, [2]int16{75, 65}, [2]int16{74, 66}, [2]int16{73, 67}, [2]int16{72, 68}, [2]int16{71, 69}, [2]int16{70, 70}, [2]int16{69, 71}, [2]int16{68, 72}, [2]int16{67, 73}, [2]int16{66, 74}, [2]int16{65, 75}, [2]int16{64, 76}, [2]int16{63, 77}, [2]int16{62, 78}, [2]int16{61, 79}, [2]int16{60, 80}, [2]int16{59, 81}, [2]int16{58, 82}, [2]int16{57, 83}, [2]int16{56, 84}, [2]int16{55, 85}, [2]int16{54, 86}, [2]int16{53, 87}, [2]int16{52, 88}, [2]int16{51, 89}, [2]int16{50, 90}, [2]int16{49, 91}, [2]int16{48, 92}, [2]int16{47, 93}, [2]int16{46, 94}, [2]int16{45, 95}, [2]int16{44, 96}, [2]int16{43, 97}, [2]int16{42, 98}, [2]int16{41, 99}, [2]int16{40, 99}, [2]int16{39, 99}, [2]int16{38, 99}, [2]int16{37, 99}, [2]int16{36, 99}, [2]int16{35, 99}, [2]int16{34, 99}, [2]int16{33, 99}, [2]int16{32, 99}, [2]int16{31, 99}, [2]int16{30, 99}, [2]int16{29, 99}, [2]int16{28, 99}, [2]int16{27, 99}, [2]int16{26, 99}, [2]int16{25, 99}, [2]int16{24, 99}, [2]int16{23, 99}, [2]int16{22, 99}, [2]int16{21, 99}, [2]int16{20, 99}, [2]int16{19, 99}, [2]int16{18, 99}, [2]int16{17, 99}, [2]int16{16, 99}, [2]int16{15, 99}, [2]int16{14, 99}, [2]int16{13, 99}, [2]int16{12, 99}, [2]int16{11, 99}, [2]int16{10, 99}, [2]int16{9, 99}, [2]int16{8, 99}, [2]int16{7, 99}, [2]int16{6, 99}, [2]int16{5, 99}, [2]int16{4, 99}, [2]int16{3, 99}, [2]int16{2, 99}, [2]int16{1, 99}, [2]int16{0, 99},}
	res = make([]Coordinates, 0, 99)
	for _, i := range coords1 {
		res = append(res, Coordinates{
			X: i[0],
			Y: i[1],
		})
	}
	case4.Result = res
	cases["case4"] = case4

	barier = make(map[Coordinates]struct{})
	for b := int16(99); b > 10; b-- {
		barier[Coordinates{
			X: b,
			Y: 20,
		}] = struct{}{}
		barier[Coordinates{
			X: b,
			Y: 70,
		}] = struct{}{}
	}
	case5 := TestBarrier{
		From:   Coordinates{99, 99},
		To:     Coordinates{99, 0},
		Result: nil,
		Fill:   barier,
	}
	coords3 := [228][2]int16{[2]int16{98, 98}, [2]int16{97, 97}, [2]int16{96, 96}, [2]int16{95, 95}, [2]int16{94, 94}, [2]int16{93, 93}, [2]int16{92, 92}, [2]int16{91, 91}, [2]int16{90, 90}, [2]int16{89, 89}, [2]int16{88, 88}, [2]int16{87, 87}, [2]int16{86, 86}, [2]int16{85, 85}, [2]int16{84, 84}, [2]int16{83, 83}, [2]int16{82, 82}, [2]int16{81, 81}, [2]int16{80, 80}, [2]int16{79, 79}, [2]int16{78, 78}, [2]int16{77, 77}, [2]int16{76, 76}, [2]int16{75, 75}, [2]int16{74, 74}, [2]int16{73, 73}, [2]int16{72, 72}, [2]int16{71, 71}, [2]int16{70, 71}, [2]int16{69, 71}, [2]int16{68, 71}, [2]int16{67, 71}, [2]int16{66, 71}, [2]int16{65, 71}, [2]int16{64, 71}, [2]int16{63, 71}, [2]int16{62, 71}, [2]int16{61, 71}, [2]int16{60, 71}, [2]int16{59, 71}, [2]int16{58, 71}, [2]int16{57, 71}, [2]int16{56, 71}, [2]int16{55, 71}, [2]int16{54, 71}, [2]int16{53, 71}, [2]int16{52, 71}, [2]int16{51, 71}, [2]int16{50, 71}, [2]int16{49, 71}, [2]int16{48, 71}, [2]int16{47, 71}, [2]int16{46, 71}, [2]int16{45, 71}, [2]int16{44, 71}, [2]int16{43, 71}, [2]int16{42, 71}, [2]int16{41, 71}, [2]int16{40, 71}, [2]int16{39, 71}, [2]int16{38, 71}, [2]int16{37, 71}, [2]int16{36, 71}, [2]int16{35, 71}, [2]int16{34, 71}, [2]int16{33, 71}, [2]int16{32, 71}, [2]int16{31, 71}, [2]int16{30, 71}, [2]int16{29, 71}, [2]int16{28, 71}, [2]int16{27, 71}, [2]int16{26, 71}, [2]int16{25, 71}, [2]int16{24, 71}, [2]int16{23, 71}, [2]int16{22, 71}, [2]int16{21, 71}, [2]int16{20, 71}, [2]int16{19, 71}, [2]int16{18, 71}, [2]int16{17, 71}, [2]int16{16, 71}, [2]int16{15, 71}, [2]int16{14, 71}, [2]int16{13, 71}, [2]int16{12, 71}, [2]int16{11, 71}, [2]int16{10, 70}, [2]int16{10, 69}, [2]int16{10, 68}, [2]int16{10, 67}, [2]int16{10, 66}, [2]int16{10, 65}, [2]int16{10, 64}, [2]int16{10, 63}, [2]int16{10, 62}, [2]int16{10, 61}, [2]int16{10, 60}, [2]int16{10, 59}, [2]int16{10, 58}, [2]int16{10, 57}, [2]int16{10, 56}, [2]int16{10, 55}, [2]int16{10, 54}, [2]int16{10, 53}, [2]int16{10, 52}, [2]int16{10, 51}, [2]int16{10, 50}, [2]int16{10, 49}, [2]int16{10, 48}, [2]int16{10, 47}, [2]int16{10, 46}, [2]int16{10, 45}, [2]int16{10, 44}, [2]int16{10, 43}, [2]int16{10, 42}, [2]int16{10, 41}, [2]int16{10, 40}, [2]int16{10, 39}, [2]int16{10, 38}, [2]int16{10, 37}, [2]int16{10, 36}, [2]int16{10, 35}, [2]int16{10, 34}, [2]int16{10, 33}, [2]int16{10, 32}, [2]int16{10, 31}, [2]int16{10, 30}, [2]int16{10, 29}, [2]int16{10, 28}, [2]int16{10, 27}, [2]int16{10, 26}, [2]int16{10, 25}, [2]int16{10, 24}, [2]int16{10, 23}, [2]int16{10, 22}, [2]int16{10, 21}, [2]int16{10, 20}, [2]int16{11, 19}, [2]int16{12, 18}, [2]int16{13, 17}, [2]int16{14, 16}, [2]int16{15, 15}, [2]int16{16, 14}, [2]int16{17, 13}, [2]int16{18, 12}, [2]int16{19, 11}, [2]int16{20, 10}, [2]int16{21, 9}, [2]int16{22, 8}, [2]int16{23, 7}, [2]int16{24, 6}, [2]int16{25, 5}, [2]int16{26, 4}, [2]int16{27, 3}, [2]int16{28, 2}, [2]int16{29, 1}, [2]int16{30, 0}, [2]int16{31, 0}, [2]int16{32, 0}, [2]int16{33, 0}, [2]int16{34, 0}, [2]int16{35, 0}, [2]int16{36, 0}, [2]int16{37, 0}, [2]int16{38, 0}, [2]int16{39, 0}, [2]int16{40, 0}, [2]int16{41, 0}, [2]int16{42, 0}, [2]int16{43, 0}, [2]int16{44, 0}, [2]int16{45, 0}, [2]int16{46, 0}, [2]int16{47, 0}, [2]int16{48, 0}, [2]int16{49, 0}, [2]int16{50, 0}, [2]int16{51, 0}, [2]int16{52, 0}, [2]int16{53, 0}, [2]int16{54, 0}, [2]int16{55, 0}, [2]int16{56, 0}, [2]int16{57, 0}, [2]int16{58, 0}, [2]int16{59, 0}, [2]int16{60, 0}, [2]int16{61, 0}, [2]int16{62, 0}, [2]int16{63, 0}, [2]int16{64, 0}, [2]int16{65, 0}, [2]int16{66, 0}, [2]int16{67, 0}, [2]int16{68, 0}, [2]int16{69, 0}, [2]int16{70, 0}, [2]int16{71, 0}, [2]int16{72, 0}, [2]int16{73, 0}, [2]int16{74, 0}, [2]int16{75, 0}, [2]int16{76, 0}, [2]int16{77, 0}, [2]int16{78, 0}, [2]int16{79, 0}, [2]int16{80, 0}, [2]int16{81, 0}, [2]int16{82, 0}, [2]int16{83, 0}, [2]int16{84, 0}, [2]int16{85, 0}, [2]int16{86, 0}, [2]int16{87, 0}, [2]int16{88, 0}, [2]int16{89, 0}, [2]int16{90, 0}, [2]int16{91, 0}, [2]int16{92, 0}, [2]int16{93, 0}, [2]int16{94, 0}, [2]int16{95, 0}, [2]int16{96, 0}, [2]int16{97, 0}, [2]int16{98, 0}, [2]int16{99, 0},}
	res = make([]Coordinates, 0, 99)
	for _, i := range coords3 {
		res = append(res, Coordinates{
			X: i[0],
			Y: i[1],
		})
	}
	case5.Result = res
	cases["case5"] = case5

	barier = make(map[Coordinates]struct{})
	for b := int16(99); b > 10; b-- {
		barier[Coordinates{
			X: 20,
			Y: b,
		}] = struct{}{}
		barier[Coordinates{
			X: 70,
			Y: b,
		}] = struct{}{}
	}
	case6 := TestBarrier{
		From:   Coordinates{99, 99},
		To:     Coordinates{0, 99},
		Result: nil,
		Fill:   barier,
	}
	coords3 = [228][2]int16{[2]int16{98, 98},[2]int16{97, 97},[2]int16{96, 96},[2]int16{95, 95},[2]int16{94, 94},[2]int16{93, 93},[2]int16{92, 92},[2]int16{91, 91},[2]int16{90, 90},[2]int16{89, 89},[2]int16{88, 88},[2]int16{87, 87},[2]int16{86, 86},[2]int16{85, 85},[2]int16{84, 84},[2]int16{83, 83},[2]int16{82, 82},[2]int16{81, 81},[2]int16{80, 80},[2]int16{79, 79},[2]int16{78, 78},[2]int16{77, 77},[2]int16{76, 76},[2]int16{75, 75},[2]int16{74, 74},[2]int16{73, 73},[2]int16{72, 72},[2]int16{71, 71},[2]int16{71, 70},[2]int16{71, 69},[2]int16{71, 68},[2]int16{71, 67},[2]int16{71, 66},[2]int16{71, 65},[2]int16{71, 64},[2]int16{71, 63},[2]int16{71, 62},[2]int16{71, 61},[2]int16{71, 60},[2]int16{71, 59},[2]int16{71, 58},[2]int16{71, 57},[2]int16{71, 56},[2]int16{71, 55},[2]int16{71, 54},[2]int16{71, 53},[2]int16{71, 52},[2]int16{71, 51},[2]int16{71, 50},[2]int16{71, 49},[2]int16{71, 48},[2]int16{71, 47},[2]int16{71, 46},[2]int16{71, 45},[2]int16{71, 44},[2]int16{71, 43},[2]int16{71, 42},[2]int16{71, 41},[2]int16{71, 40},[2]int16{71, 39},[2]int16{71, 38},[2]int16{71, 37},[2]int16{71, 36},[2]int16{71, 35},[2]int16{71, 34},[2]int16{71, 33},[2]int16{71, 32},[2]int16{71, 31},[2]int16{71, 30},[2]int16{71, 29},[2]int16{71, 28},[2]int16{71, 27},[2]int16{71, 26},[2]int16{71, 25},[2]int16{71, 24},[2]int16{71, 23},[2]int16{71, 22},[2]int16{71, 21},[2]int16{71, 20},[2]int16{71, 19},[2]int16{71, 18},[2]int16{71, 17},[2]int16{71, 16},[2]int16{71, 15},[2]int16{71, 14},[2]int16{71, 13},[2]int16{71, 12},[2]int16{71, 11},[2]int16{70, 10},[2]int16{69, 10},[2]int16{68, 10},[2]int16{67, 10},[2]int16{66, 10},[2]int16{65, 10},[2]int16{64, 10},[2]int16{63, 10},[2]int16{62, 10},[2]int16{61, 10},[2]int16{60, 10},[2]int16{59, 10},[2]int16{58, 10},[2]int16{57, 10},[2]int16{56, 10},[2]int16{55, 10},[2]int16{54, 10},[2]int16{53, 10},[2]int16{52, 10},[2]int16{51, 10},[2]int16{50, 10},[2]int16{49, 10},[2]int16{48, 10},[2]int16{47, 10},[2]int16{46, 10},[2]int16{45, 10},[2]int16{44, 10},[2]int16{43, 10},[2]int16{42, 10},[2]int16{41, 10},[2]int16{40, 10},[2]int16{39, 10},[2]int16{38, 10},[2]int16{37, 10},[2]int16{36, 10},[2]int16{35, 10},[2]int16{34, 10},[2]int16{33, 10},[2]int16{32, 10},[2]int16{31, 10},[2]int16{30, 10},[2]int16{29, 10},[2]int16{28, 10},[2]int16{27, 10},[2]int16{26, 10},[2]int16{25, 10},[2]int16{24, 10},[2]int16{23, 10},[2]int16{22, 10},[2]int16{21, 10},[2]int16{20, 10},[2]int16{19, 11},[2]int16{18, 12},[2]int16{17, 13},[2]int16{16, 14},[2]int16{15, 15},[2]int16{14, 16},[2]int16{13, 17},[2]int16{12, 18},[2]int16{11, 19},[2]int16{10, 20},[2]int16{9, 21},[2]int16{8, 22},[2]int16{7, 23},[2]int16{6, 24},[2]int16{5, 25},[2]int16{4, 26},[2]int16{3, 27},[2]int16{2, 28},[2]int16{1, 29},[2]int16{0, 30},[2]int16{0, 31},[2]int16{0, 32},[2]int16{0, 33},[2]int16{0, 34},[2]int16{0, 35},[2]int16{0, 36},[2]int16{0, 37},[2]int16{0, 38},[2]int16{0, 39},[2]int16{0, 40},[2]int16{0, 41},[2]int16{0, 42},[2]int16{0, 43},[2]int16{0, 44},[2]int16{0, 45},[2]int16{0, 46},[2]int16{0, 47},[2]int16{0, 48},[2]int16{0, 49},[2]int16{0, 50},[2]int16{0, 51},[2]int16{0, 52},[2]int16{0, 53},[2]int16{0, 54},[2]int16{0, 55},[2]int16{0, 56},[2]int16{0, 57},[2]int16{0, 58},[2]int16{0, 59},[2]int16{0, 60},[2]int16{0, 61},[2]int16{0, 62},[2]int16{0, 63},[2]int16{0, 64},[2]int16{0, 65},[2]int16{0, 66},[2]int16{0, 67},[2]int16{0, 68},[2]int16{0, 69},[2]int16{0, 70},[2]int16{0, 71},[2]int16{0, 72},[2]int16{0, 73},[2]int16{0, 74},[2]int16{0, 75},[2]int16{0, 76},[2]int16{0, 77},[2]int16{0, 78},[2]int16{0, 79},[2]int16{0, 80},[2]int16{0, 81},[2]int16{0, 82},[2]int16{0, 83},[2]int16{0, 84},[2]int16{0, 85},[2]int16{0, 86},[2]int16{0, 87},[2]int16{0, 88},[2]int16{0, 89},[2]int16{0, 90},[2]int16{0, 91},[2]int16{0, 92},[2]int16{0, 93},[2]int16{0, 94},[2]int16{0, 95},[2]int16{0, 96},[2]int16{0, 97},[2]int16{0, 98},[2]int16{0, 99},}
	res = make([]Coordinates, 0, 99)
	for _, i := range coords3 {
		res = append(res, Coordinates{
			X: i[0],
			Y: i[1],
		})
	}
	case6.Result = res
	cases["case6"] = case6

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

		result := reflect.DeepEqual(testCase.Result, astar.FindPath())
		if !result {
			t.Errorf("slice not equal in %s", i)
			fmt.Println(testCase.Result)
			fmt.Println(astar.FindPath())
		}
	}
}
