package exercise2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneRound(t *testing.T) {
	result := HowManyRounds("13:45", "14:00")

	assert.Equal(t, 1, result)
}

func TestOneRoundOvernight(t *testing.T) {
	result := HowManyRounds("23:56", "00:15")

	assert.Equal(t, 1, result)
}

func TestFourRound(t *testing.T) {
	result := HowManyRounds("23:00", "00:3")

	assert.Equal(t, 4, result)
}

func TestNoRoundsProblematicCleaned(t *testing.T) {
	result := HowManyRounds("23:16", "23:17")

	assert.Equal(t, 0, result)
}

func TestNoRoundsOvernight(t *testing.T) {
	result := HowManyRounds("23:55", "0:5")

	assert.Equal(t, 0, result)
}

func TestNoRounds(t *testing.T) {
	result := HowManyRounds("0:0", "0:5")

	assert.Equal(t, 0, result)
}

func TestMaxRounds(t *testing.T) {
	result1 := HowManyRounds("20:01", "20:0")
	result2 := HowManyRounds("20:14", "20:0")

	assert.Equal(t, 95, result1, result2)
}
