package main

import (
	"fmt"
)

func main() {

	doSomething()
}
func doSomething() {
	fmt.Println("Doing Something")
	val1 := 3
	val2 := 5
	val3 := 6
	// sum := addValues(val1, val2)
	sum, count, avg := addAllValues(val1, val2, val3)

	// fmt.Printf("the sum of %v and %v id: %v\n", val1, val2, sum)
	fmt.Printf("the sum is: %v \n", sum)
	fmt.Printf("the count is: %v \n", count)
	fmt.Printf("the average is: %.2f \n", avg)
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int, float64) {
	sum := 0 
	for _, v := range values {
		sum += v
	}
	count := len(values)
	avg := float64(sum) / float64(count)
	return sum, count, avg
}