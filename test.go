package main

import (
	"github.com/JimTornvall/HA-TDD-MarsRover/rover"
	"testing"
)

func TestNewRover(t *testing.T) {
	x := 1
	y := 2
	pos := rover.NewPosition(x, y)
	dir := rover.Direction("N")
	r := rover.NewRover(x, y, dir)
	if r.GetPosition() != pos {
		t.Errorf("Expected pos to be %v, got %v", pos, r.GetPosition())
	}
	if r.GetDirection() != dir {
		t.Errorf("Expected dir to be %v, got %v", dir, r.GetDirection())
	}
}
