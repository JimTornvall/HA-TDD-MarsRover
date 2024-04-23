package main

import (
	"github.com/JimTornvall/HA-TDD-MarsRover/rover"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRover(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	pos := rover.NewPosition(x, y)
	dir := rover.Direction("N")

	// Act
	r, _ := rover.NewRover(x, y, dir)

	// Assert
	assert.Equal(t, pos, r.GetPosition(), "Expected pos to be %v, got %v", pos, r.GetPosition())
	assert.Equal(t, dir, r.GetDirection(), "Expected dir to be %v, got %v", dir, r.GetDirection())
}

func TestMoveForward(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	pos := rover.NewPosition(x, y)

	expectedX := 1
	expectedY := 1
	expectedPos := rover.NewPosition(expectedX, expectedY)

	dir := rover.Direction("N")
	r, _ := rover.NewRover(x, y, dir)

	// Act
	_ = r.MoveForward()

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", pos, r.GetPosition())
	assert.Equal(t, dir, r.GetDirection(), "Expected dir to be %v, got %v", dir, r.GetDirection())
}

func TestMoveBackward(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	pos := rover.NewPosition(x, y)

	expectedX := 1
	expectedY := 3
	expectedPos := rover.NewPosition(expectedX, expectedY)

	dir := rover.Direction("N")
	r, _ := rover.NewRover(x, y, dir)

	// Act
	_ = r.MoveBackward()

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", pos, r.GetPosition())
	assert.Equal(t, dir, r.GetDirection(), "Expected dir to be %v, got %v", dir, r.GetDirection())
}
