package main

import (
	"fmt"
)

func main() {
	maps := map[int]string{ // declare map
		0: "0",
		1: "1",
	}

	var maps2 = make(map[int]string) // only define
	fmt.Println(maps2)

	for i, v := range maps { // loop
		fmt.Println(i, ":", v)
	}
	fmt.Println(maps[0]) // retrieve
	delete(maps, 0)      // delete element by key
	fmt.Println(maps)
}
