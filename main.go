package main

import (
	"exercises/exercise1"
	"exercises/exercise2"
	"exercises/exercise3"
	"fmt"
)

func main() {
	fmt.Println(exercise1.HowManyBananaInString("NAABXXAN"))
	fmt.Println(exercise2.HowManyRounds("20:01", "20:0"))
	fmt.Println(exercise3.WhichOperation("nice", "nicer"))
}
