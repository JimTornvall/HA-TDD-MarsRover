package rover

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
		r.pos.x++
	case "S":
		r.pos.y++
	case "W":
		r.pos.x--
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
		r.pos.x--
	case "S":
		r.pos.y--
	case "W":
		r.pos.x++
	}
	return nil
}

func (r *Rover) TurnLeft() {
	d := r.GetDirection()
	switch d {
	case "N":
		r.dir = "W"
	case "W":
		r.dir = "S"
	case "S":
		r.dir = "E"
	case "E":
		r.dir = "N"
	}
}

func (r *Rover) TurnRight() {
	d := r.GetDirection()
	switch d {
	case "N":
		r.dir = "E"
	case "E":
		r.dir = "N"
	case "S":
		r.dir = "W"
	case "W":
		r.dir = "N"
	}

}
