package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRover(t *testing.T) {

	// Arrange
	x := 1
	y := 2
	pos := Position{x, y}
	dir := Direction("N")

	// Act
	r, _ := NewRover(x, y, dir)

	// Assert
	assert.Equal(t, pos, r.pos, "Expected pos to be %v, got %v", pos, r.pos)
	assert.Equal(t, dir, r.dir, "Expected dir to be %v, got %v", dir, r.dir)
}
