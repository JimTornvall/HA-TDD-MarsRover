package rover

import "fmt"

// X needs to be an uneven number to make the grid symmetrical
const gridMaxX = 5
const gridMaxY = 5

type Grid struct {
	grid [gridMaxX][gridMaxY]bool
}

func NewGrid() Grid {
	return Grid{}
}

func (grid *Grid) SetObstacle(x int, y int) {
	grid.grid[x][y] = true
}
func (grid *Grid) IsObstacle(x int, y int) bool {
	return grid.grid[x][y]
}
func (grid *Grid) ClearObstacle(x int, y int) {
	grid.grid[x][y] = false
}

func (grid *Grid) IsValidPosition(x int, y int) bool {
	return !(x >= 0 && x <= gridMaxX && y >= 0 && y <= gridMaxY)
}

// returns false if the rover can't move to the new position
func (rover *Rover) moveNorth() error {
	x := rover.pos.x
	y := rover.pos.y
	if rover.grid.IsValidPosition(x, y) {
		return fmt.Errorf("invalid position: %v, %v", x, y)
	}
	if y == 0 {
		if x < (gridMaxX+1)/2 {
			newX := x + (gridMaxX+1)/2
			newY := y
			if rover.grid.IsObstacle(newX, newY) {
				return fmt.Errorf("obstacle at: %v, %v", newX, newY)
			}
			rover.pos.x = newX
			rover.pos.y = newY
			rover.dir.Switch()
			return nil
		} else {
			newX := x - (gridMaxX+1)/2
			newY := y
			if rover.grid.IsObstacle(newX, newY) {
				return fmt.Errorf("obstacle at: %v, %v", newX, newY)
			}
			rover.pos.x = newX
			rover.pos.y = newY
			rover.dir.Switch()
			return nil
		}
	}
	rover.pos.y--
	return nil
}

func (rover *Rover) moveEast(x int, y int) (int, int) {
	if x == gridMaxX-1 {
		return 0, y
	}
	return x + 1, y

}

func (rover *Rover) moveSouth(x int, y int) (int, int) {
	if y == gridMaxY-1 {
		return x, 0
	}
	return x, y + 1
}

func (rover *Rover) moveWest(x int, y int) (int, int) {
	if x == 0 {
		return gridMaxX - 1, y
	}
	return x - 1, y
}
