package rover

import (
	"fmt"
)

type Direction rune

//    N
//  W   E
//    S

func NewDirection(dir rune) (Direction, error) {
	if dir != 'N' && dir != 'E' && dir != 'S' && dir != 'W' {
		return Direction(0), fmt.Errorf("invalid direction: %v", dir)
	}
	return Direction(dir), nil
}

func (d *Direction) Switch() {
	switch *d {
	case 'N':
		*d = 'S'
	case 'S':
		*d = 'N'
	case 'E':
		*d = 'W'
	case 'W':
		*d = 'E'
	}
}

func (d Direction) String() string {
	return string(d)
}
