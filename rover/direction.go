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
