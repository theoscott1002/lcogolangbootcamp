package main

import (
	"fmt"
)

func main() {
	var lifeBooster float64 = 84.5
	lifeBoosterRef := &lifeBooster

	fmt.Println(lifeBooster)
	fmt.Println(*lifeBoosterRef)

}
