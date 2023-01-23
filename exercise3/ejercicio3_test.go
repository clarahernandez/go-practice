package exercise3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	solution := WhichOperation("nice", "nicer")

	assert.Equal(t, solution, "ADD r")
}

func TestJoin(t *testing.T) {
	solution := WhichOperation("meet", "met")

	assert.Equal(t, solution, "JOIN e")
}

func TestSwap(t *testing.T) {
	solution := WhichOperation("late", "tale")

	assert.Equal(t, solution, "SWAP l t")
}

func TestNothing(t *testing.T) {
	solution := WhichOperation("hello", "hello")

	assert.Equal(t, solution, "NOTHING")
}
