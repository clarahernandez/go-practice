package exercise1

import (
	"math"
	"strings"
)

func HowManyBananaInString(S string) int {
	bCounter := 0
	aCounter := 0
	nCounter := 0

	upperS := strings.ToUpper(S)

	for i := 0; i < len(upperS); i++ {
		if upperS[i] == 'B' {
			bCounter++
		}
		if upperS[i] == 'A' {
			aCounter++
		}
		if upperS[i] == 'N' {
			nCounter++
		}
	}

	return int(min(float64(bCounter), float64(aCounter/3), float64(nCounter/2)))
}

func min(x float64, y float64, z float64) float64 {
	return math.Min(math.Min(x, y), z)
}
