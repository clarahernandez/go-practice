package exercise1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneBanana(t *testing.T) {
	result := HowManyBananaInString("NAABXXAN")

	assert.Equal(t, int(1), result)
}

func TestTwoBanana(t *testing.T) {
	result := HowManyBananaInString("BXBAXAXNNXAXAXNNXAXA")

	assert.Equal(t, 2, result)
}

func TestNoneAllB(t *testing.T) {
	result := HowManyBananaInString("ANANA")

	assert.Equal(t, 0, result)
}

func TestNoneAllA(t *testing.T) {
	result := HowManyBananaInString("BANAN")

	assert.Equal(t, 0, result)
}

func TestNoneAllN(t *testing.T) {
	result := HowManyBananaInString("BANAA")

	assert.Equal(t, 0, result)
}

func TestBananaLowerCase(t *testing.T) {
	result := HowManyBananaInString("banana")

	assert.Equal(t, 1, result)
}

func TestBananaReverse(t *testing.T) {
	result := HowManyBananaInString("ANANAB")

	assert.Equal(t, 1, result)
}

func TestBananaAlmost3WithA(t *testing.T) {
	result := HowManyBananaInString("BBBAAANNNAAANNNAA")

	assert.Equal(t, 2, result)
}
