package main

import "fmt"

func main() {
	i1, i2, i3 := 52, 35, 33
	integerSum := i1 + i2 + i3
	fmt.Println("Integer sum:", integerSum)

	f1, f2, f3 := 22.03, 41.66, 20.3
	FloatSum := f1 + f2 + f3
	fmt.Println("Float sum:", FloatSum)

	total := float64(i1) + f1
	fmt.Println("Result:", total)

	product := float64(i2) * f3
	fmt.Println("Product:", product)

	division := float64(i3) / f1
	fmt.Println("Division:", division)

}
