package rover

type Rover struct {
	pos  Position
	dir  Direction
	grid Grid
}

func NewRover(x int, y int, dir Direction) (Rover, error) {
	// TODO: validate x, y, dir

	pos := Position{x, y}
	return Rover{pos, dir, NewGrid()}, nil
}

func (rover *Rover) GetPosition() Position {
	// TODO: should return x and y as int

	return rover.pos
}

func (rover *Rover) GetDirection() Direction {
	return rover.dir
}

func (rover *Rover) MoveForward() error {
	// TODO: should return error if rover found an obstacle
	d := rover.GetDirection()
	var err error
	switch d {
	case 'N':
		//rover.pos.y--
		err = rover.moveNorth()
	case 'E':
		rover.pos.x++
	case 'S':
		rover.pos.y++
	case 'W':
		rover.pos.x--
	}
	return err
}

func (rover *Rover) MoveBackward() error {
	// TODO: should return error if rover found an obstacle
	var err error
	d := rover.GetDirection()
	switch d {
	case 'N':
		rover.pos.y++
	case 'E':
		rover.pos.x--
	case 'S':
		//rover.pos.y--
		err = rover.moveNorth()
	case 'W':
		rover.pos.x++
	}
	return err
}

func (rover *Rover) TurnLeft() {
	d := rover.GetDirection()
	switch d {
	case 'N':
		rover.dir = 'W'
	case 'W':
		rover.dir = 'S'
	case 'S':
		rover.dir = 'E'
	case 'E':
		rover.dir = 'N'
	}
}

func (rover *Rover) TurnRight() {
	d := rover.GetDirection()
	switch d {
	case 'N':
		rover.dir = 'E'
	case 'E':
		rover.dir = 'S'
	case 'S':
		rover.dir = 'W'
	case 'W':
		rover.dir = 'N'
	}
}

func (rover *Rover) Run(commands string) error {
	for _, c := range commands {
		switch c {
		case 'F':
			err := rover.MoveForward()
			if err != nil {
				return err
			}
		case 'B':
			err := rover.MoveBackward()
			if err != nil {
				return err
			}
		case 'L':
			rover.TurnLeft()
		case 'R':
			rover.TurnRight()
		}
	}
	return nil
}
