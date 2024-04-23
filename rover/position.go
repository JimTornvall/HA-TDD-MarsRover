package rover

import "fmt"

type Position struct {
	x int
	y int
}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func (p Position) String() string {
	return fmt.Sprint(p.x, p.y)
}
