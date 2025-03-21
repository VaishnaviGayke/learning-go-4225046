package main

import (
	"fmt"
	"time"
)

func main() {

	timeofGO := time.Date(2009, time.November, 13, 22, 15, 0, 0, time.UTC)
	fmt.Printf("Dates and times of GO: %v\n", timeofGO)

	nowtime := time.Now()
	fmt.Printf("Current time is: %v\n", nowtime)
	fmt.Printf("The object type is : %T\n", timeofGO)

	fmt.Printf(nowtime.Format(time.ANSIC) + "\n")

	tom := nowtime.AddDate(0,0,1)
	fmt.Printf(tom.Format(time.ANSIC) + "\n")

	formatbyown := "Mon 2006-02-01"
	fmt.Printf(tom.Format(formatbyown) + "\n")

}
