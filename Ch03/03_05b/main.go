package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["CA"] = "California"
	fmt.Println(states)

	// calii := states["CA"]
	fmt.Println(states["CA"])

	delete(states, "WA")
	fmt.Println(states)

	states["TX"] = "Texas"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k,v)
	}

	keys := make([]string, len(states))
	i :=0
	for k := range states{
		keys[i] = k
		i++
	}

	fmt.Println(keys)

	sort.Strings(keys)
	fmt.Println("sorted keys slice")
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
