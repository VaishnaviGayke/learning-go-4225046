package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is a slice
	// var colors = []string{"Red", "Green", "Blue"}

	var colors = make([]string,0,3 )
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	colors = append(colors, "Yellow", "Purple", "White")
	fmt.Println(colors)

	colors =  remove(colors, 0)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

func remove(slice []string, index int) []string{
	return append(slice[:index], slice[index+1:]...)
}
