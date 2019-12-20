package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 3, 5, 6} // declare int slice
	slice = append(slice, 8)
	fmt.Println(cap(slice)) // 8
	fmt.Println(len(slice)) // 5
	fmt.Println(slice[:3])  // [1 3 5]

	slice2 := make([]int, 10) // define with capacity
	fmt.Println(cap(slice2))  // 10

	// define 2 dim slice
	rows := 8
	cols := 8
	slice3 := make([][]int, rows)
	for i := range slice3 {
		slice3[i] = make([]int, cols)
	}
}
