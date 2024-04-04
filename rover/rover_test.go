package rover

import "testing"

func TestNewRover(t *testing.T) {
	x := 1
	y := 2
	pos := Position{x, y}
	dir := Direction("N")
	r := NewRover(x, y, dir)
	if r.pos != pos {
		t.Errorf("Expected pos to be %v, got %v", pos, r.pos)
	}
	if r.dir != dir {
		t.Errorf("Expected dir to be %v, got %v", dir, r.dir)
	}
}
