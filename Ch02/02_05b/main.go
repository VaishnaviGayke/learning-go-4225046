package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum is updated as %v\n", sum)

	fmt.Println("The value of PI is", math.Pi)

	circleRadius := 20.33
	circum := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference is: %.2f\n", circum)
}
