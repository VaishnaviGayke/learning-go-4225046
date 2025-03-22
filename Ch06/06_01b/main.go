package main

import (
	"fmt"
	"io"
	"os"
	
)

func main() {
	filename := "./fromString.txt"
	file, errr := os.Create(filename)
	defer file.Close()
	
	checkError(errr)

	length, errr := io.WriteString(file, "Hello from GO!!!")
	fmt.Printf("Wrote a file with %v characters\n", length)
	readFile(filename)
}

func checkError(errr error){
	if errr != nil {
		panic(errr)
	}
}

func readFile(filename string){
	data, errr := os.ReadFile(filename)
	checkError(errr)
	fmt.Println("Text read from readFile function:\n", string(data))
}