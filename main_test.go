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
	dir := rover.Direction('N')

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

	expectedX := 1
	expectedY := 1
	expectedPos := rover.NewPosition(expectedX, expectedY)

	dir := rover.Direction('N')
	r, _ := rover.NewRover(x, y, dir)

	// Act
	_ = r.MoveForward()

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
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

	dir := rover.Direction('N')
	r, _ := rover.NewRover(x, y, dir)

	// Act
	_ = r.MoveBackward()

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", pos, r.GetPosition())
	assert.Equal(t, dir, r.GetDirection(), "Expected dir to be %v, got %v", dir, r.GetDirection())
}

func TestTurnLeft(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	pos := rover.NewPosition(x, y)

	dir := rover.Direction('N')
	expectedDir := rover.Direction('W')
	r, _ := rover.NewRover(x, y, dir)

	// Act
	r.TurnLeft()

	// Assert
	assert.Equal(t, pos, r.GetPosition(), "Expected pos to be %v, got %v", pos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
}

func TestTurnRight(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	pos := rover.NewPosition(x, y)

	dir := rover.Direction('N')
	expectedDir := rover.Direction('E')
	r, _ := rover.NewRover(x, y, dir)

	// Act
	r.TurnRight()

	// Assert
	assert.Equal(t, pos, r.GetPosition(), "Expected pos to be %v, got %v", pos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
}

func TestRunCommands(t *testing.T) {
	// Arrange
	x := 2 //2 3 3
	y := 1 //1 0 1
	dir, _ := rover.NewDirection('N')

	expectedDir := rover.Direction('N')
	expectedPos := rover.NewPosition(3, 1)
	r, _ := rover.NewRover(x, y, dir)

	commands := "FRFLB"

	// Act
	_ = r.Run(commands)

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
}

func TestRunCommands2(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	dir, _ := rover.NewDirection('N')

	expectedDir := rover.Direction('N')
	expectedPos := rover.NewPosition(1, 2)
	r, _ := rover.NewRover(x, y, dir)

	commands := "FBFBLRLR"

	// Act
	_ = r.Run(commands)

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
}

func TestRunEdgeDetectionNorth(t *testing.T) {
	// Arrange
	x := 0 // 0 3 3
	y := 1 // 0 0 1
	dir, _ := rover.NewDirection('N')

	expectedPos := rover.NewPosition(3, 1)
	var expectedDir rover.Direction = 'S'
	r, _ := rover.NewRover(x, y, dir)

	commands := "FFF"

	// Act
	err := r.Run(commands)

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
	assert.Nil(t, err, "Expected error to be nil")
}

func TestRunEdgeDetectionSouth(t *testing.T) {
	// Arrange
	x := 2                  //
	y := rover.GridMaxY - 1 //
	dir, _ := rover.NewDirection('N')

	expectedPos := rover.NewPosition(5, 4)
	var expectedDir rover.Direction = 'S'
	r, _ := rover.NewRover(x, y, dir)

	commands := "BBB"

	// Act
	err := r.Run(commands)

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
	assert.Nil(t, err, "Expected error to be nil")
}

func TestRunEdgeDetectionEast(t *testing.T) {
	// Arrange
	x := rover.GridMaxX - 1
	y := 2
	dir, _ := rover.NewDirection('E')

	expectedPos := rover.NewPosition(1, 2)
	var expectedDir rover.Direction = 'E'
	r, _ := rover.NewRover(x, y, dir)

	commands := "FFF"

	// Act
	err := r.Run(commands)

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
	assert.Nil(t, err, "Expected error to be nil")
}

func TestRunEdgeDetectionWest(t *testing.T) {
	// Arrange
	x := 1
	y := 2
	dir, _ := rover.NewDirection('E')

	expectedPos := rover.NewPosition(4, 2)
	var expectedDir rover.Direction = 'E'
	r, _ := rover.NewRover(x, y, dir)

	commands := "BBB"

	// Act
	err := r.Run(commands)

	// Assert
	assert.Equal(t, expectedPos, r.GetPosition(), "Expected pos to be %v, got %v", expectedPos, r.GetPosition())
	assert.Equal(t, expectedDir, r.GetDirection(), "Expected dir to be %v, got %v", expectedDir, r.GetDirection())
	assert.Nil(t, err, "Expected error to be nil")
}
