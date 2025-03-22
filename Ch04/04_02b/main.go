package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := uint(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	
	switch dayNumber = 2; dayNumber{
	case 1: 
		result = "It's a Monday"
	case 2: 
		result = "It's a Tuesday"
	case 3: 
		result = "It's a Wednesday"
	case 4: 
		result = "It's a Thursday"
	case 5: 
		result = "It's a Friday"
	default : 
		result = "It's the Weekend"
	}
	fmt.Println(result)


	x:= 0
	switch{
	case x<0:
		result = "Less than Zero"
	case x==0:
		result = "Equal to Zero"
	default:
		result = "Greater than Zero"
	}
	fmt.Println(result)


}
