package main

import (
	"fmt"
)

// array: fixed size
// if add elements over than the size, copy and paste
func main() {
	var arr = [5]int{1, 2, 4, 5, 6} // declare int array

	for i := 0; i < len(arr); i++ { // len(): find size of array
		fmt.Println(arr[i]) // arr
	}

	var isTrue bool
	for i, v := range arr { // index, value
		isTrue = (v == arr[i])
		fmt.Println(isTrue) // true
	}

	var array2d [8][8]int
	array2d[3][6] = 18
	fmt.Println(array2d)
}
