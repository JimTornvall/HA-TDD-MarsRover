package rover

// ------------------------------
// Position
// TODO: should be private

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

// ------------------------------
// Direction
//TODO: should be private
//TODO: should have NewDirection function

type Direction string

//    N
//  E   W
//    S

// ------------------------------
// Rover

type Rover struct {
	pos Position
	dir Direction
}

func NewRover(x int, y int, dir Direction) (Rover, error) {
	// TODO: validate x, y, dir

	pos := Position{x, y}
	return Rover{pos, dir}, nil
}

func (r *Rover) GetPosition() Position {
	// TODO: should return x and y as int

	return r.pos
}

func (r *Rover) GetDirection() Direction {
	return r.dir
}

func (r *Rover) MoveForward() error {
	// TODO: should return error if rover found an obstacle
	d := r.GetDirection()
	switch d {
	case "N":
		r.pos.y--
	case "E":
		r.pos.x--
	case "S":
		r.pos.y++
	case "W":
		r.pos.x++
	}
	return nil
}

func (r *Rover) MoveBackward() error {
	// TODO: should return error if rover found an obstacle
	d := r.GetDirection()
	switch d {
	case "N":
		r.pos.y++
	case "E":
		r.pos.x++
	case "S":
		r.pos.y--
	case "W":
		r.pos.x--
	}
	return nil
}

// ------------------------------
