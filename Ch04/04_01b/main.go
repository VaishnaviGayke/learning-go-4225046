package main

import (
	"fmt"
)

func main() {
	// theAns := 00
	var result string

	if theAns := 42; theAns <0 {
		result = "Less than zero"	
	}else if theAns == 0 {
		result = "Equal to zero"
	}else{
		result = "Greater than zero"
	}

	fmt.Println(result)
}
