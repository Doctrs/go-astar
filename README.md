# A* Golang

The A* pathfinding algorithm is a pathfinding algorithm noted for its performance and accuracy and is commonly used in game development.

### Usage

Create map
```go
worldMap :=  Map{
    Width:  100,
    Height: 100,
}
```
Add blocking cells
```go
blockedCells = make(map[Coordinates]struct{})

blockedCells[Coordinates{
	X: 10,
	Y: 10,
}] = struct{}{}
blockedCells[Coordinates{
	X: 50,
	Y: 50,
}] = struct{}{}

worldMap.Blocked = blockedCells
```
Add From(start) position and To(goal) position
```go
worldMap.From = Coordinates{0, 0}
worldMap.To = Coordinates{99, 99}
```
Find path
```go
var path []Coordinates
path = worldMap.FindPath()
```