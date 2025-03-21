package main

import (
	"fmt"
)

func main() {
	aninte := 42 
	var p *int = &aninte

	if p==nil {
		fmt.Println("p is nil ")
	}else{
		fmt.Println("Value of p:", *p)
	}

	value1 := 44.3
	pointer1 := &value1
	*pointer1 = *pointer1/20
	fmt.Println("*pointer1 is:", *pointer1)
	fmt.Println("Original value1 is:", value1)
	fmt.Println("Pointer1 is:", *pointer1)

}
