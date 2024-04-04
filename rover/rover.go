package rover

type Position struct {
	x int
	y int
}

type Direction string

type Rover struct {
	pos Position
	dir Direction
}
