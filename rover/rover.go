package rover

// ------------------------------
// Position

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

// ------------------------------
// Direction

type Direction string

// ------------------------------
// Rover

type Rover struct {
	pos Position
	dir Direction
}

func NewRover(x int, y int, dir Direction) Rover {
	pos := Position{x, y}
	return Rover{pos, dir}
}

func (r *Rover) GetPosition() Position {
	return r.pos
}

func (r *Rover) GetDirection() Direction {
	return r.dir
}

// ------------------------------
