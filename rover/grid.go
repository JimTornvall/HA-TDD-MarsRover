package rover

import "fmt"

// GridMaxX needs to be an uneven number to make the Grid symmetrical
const GridMaxX = 5
const GridMaxY = 5

type Grid struct {
	grid [GridMaxX + 1][GridMaxY + 1]bool
}

func NewGrid() Grid {
	return Grid{}
}

func (grid *Grid) SetObstacle(x int, y int) {
	grid.grid[x][y] = true
}
func (grid *Grid) IsObstacle(x int, y int) bool {
	if !grid.IsValidPosition(x, y) {
		return false
	}
	return grid.grid[x][y]
}
func (grid *Grid) ClearObstacle(x int, y int) {
	grid.grid[x][y] = false
}

func (grid *Grid) IsValidPosition(x int, y int) bool {
	return x >= 0 && x <= GridMaxX && y >= 0 && y <= GridMaxY
}

func (rover *Rover) moveNorth() error {
	x := rover.pos.x
	y := rover.pos.y
	if y == 0 {
		if x < (GridMaxX+1)/2 {
			newX := x + (GridMaxX+1)/2
			newY := y
			if rover.Grid.IsObstacle(newX, newY) {
				return fmt.Errorf("obstacle at: %v, %v", newX, newY)
			}
			rover.pos.x = newX
			rover.pos.y = newY
			rover.dir.Switch()
			return nil
		} else {
			newX := x - (GridMaxX+1)/2
			newY := y
			if rover.Grid.IsObstacle(newX, newY) {
				return fmt.Errorf("obstacle at: %v, %v", newX, newY)
			}
			rover.pos.x = newX
			rover.pos.y = newY
			rover.dir.Switch()
			return nil
		}
	}
	newY := y - 1
	if rover.Grid.IsObstacle(x, newY) {
		return fmt.Errorf("obstacle at: %v, %v", x, newY)
	}
	rover.pos.y--
	return nil
}

func (rover *Rover) moveEast() error {
	x := rover.pos.x
	if x == GridMaxX {
		newX := 0
		if rover.Grid.IsObstacle(newX, rover.pos.y) {
			return fmt.Errorf("obstacle at: %v, %v", newX, rover.pos.y)
		}
		rover.pos.x = newX
		return nil
	}
	newX := x + 1
	if rover.Grid.IsObstacle(newX, rover.pos.y) {
		return fmt.Errorf("obstacle at: %v, %v", newX, rover.pos.y)
	}
	rover.pos.x++
	return nil
}

func (rover *Rover) moveSouth() error {
	x := rover.pos.x
	y := rover.pos.y

	if y == GridMaxY {
		if x < (GridMaxX+1)/2 {
			newX := x + (GridMaxX+1)/2
			newY := y
			if rover.Grid.IsObstacle(newX, newY) {
				return fmt.Errorf("obstacle at: %v, %v", newX, newY)
			}
			rover.pos.x = newX
			rover.pos.y = newY
			rover.dir.Switch()
			return nil
		} else {
			newX := x - (GridMaxX+1)/2
			newY := y
			if rover.Grid.IsObstacle(newX, newY) {
				return fmt.Errorf("obstacle at: %v, %v", newX, newY)
			}
			rover.pos.x = newX
			rover.pos.y = newY
			rover.dir.Switch()
			return nil
		}

	}
	newY := y + 1
	if rover.Grid.IsObstacle(x, newY) {
		return fmt.Errorf("obstacle at: %v, %v", x, newY)
	}
	rover.pos.y++
	return nil
}

func (rover *Rover) moveWest() error {
	x := rover.pos.x
	if x == 0 {
		newX := GridMaxX
		if rover.Grid.IsObstacle(newX, rover.pos.y) {
			return fmt.Errorf("obstacle at: %v, %v", newX, rover.pos.y)
		}
		rover.pos.x = newX
		return nil
	}
	newX := x - 1
	if rover.Grid.IsObstacle(newX, rover.pos.y) {
		return fmt.Errorf("obstacle at: %v, %v", newX, rover.pos.y)
	}
	rover.pos.x--
	return nil
}
