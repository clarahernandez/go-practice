package main

import (
	"ejercicios/exercise1"
	"ejercicios/exercise2"
	"ejercicios/exercise3"
	"fmt"
)

func main() {
	fmt.Println(exercise1.HowManyBananaInString("NAABXXAN"))
	fmt.Println(exercise2.HowManyRounds("20:01", "20:0"))
	fmt.Println(exercise3.WhichOperation("nice", "nicer"))
}
