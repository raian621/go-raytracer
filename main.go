package main

import "fmt"

func main() {
	fmt.Println("Hello")
	arr := []int{1, 2, 3, 4}
	arr2 := make([]int, 4)
	copy(arr2, arr)
	arr2[0] = 0

	fmt.Println(arr)
	fmt.Println(arr2)
}
