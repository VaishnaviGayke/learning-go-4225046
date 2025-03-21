package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println(str1, str2, str3)
	StringLen, err := fmt.Println("The value is", aNumber)
	if err == nil {
		fmt.Println("String length:", StringLen)
	}
	fmt.Printf("Value of aNumber: %v \n", aNumber)

	fmt.Printf("Data Type of aNumber: %T \n", aNumber)
}
