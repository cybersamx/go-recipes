package main

import "fmt"

func squareArrayVal(array [3]int) {
	for i, v := range array {
		array[i] = v * v
	}
}
func squareArrayPtr(array *[3]int) {
	for i, v := range *array {
		array[i] = v * v
	}
}
func squareSliceVal(slice []int) {
	for i, v := range slice {
		slice[i] = v * v
	}
}
func squareSlicePtr(slice *[]int) {
	for i, v := range *slice {
		(*slice)[i] = v * v
	}
}
func main() {
	array := [3]int{1, 2, 3}

	fmt.Println("--- Array ---")
	fmt.Println(array)

	// Pass the array as a value. We can't modify the original array.
	squareArrayVal(array)
	fmt.Println(array)

	// Pass the array as a pointer. We can modify the original array.
	squareArrayPtr(&array)
	fmt.Println(array)

	slice := []int{1, 2, 3}

	fmt.Println("--- Slice ---")
	fmt.Println(slice)

	// Pass the slice as a value. We can modify the original slice.
	squareSliceVal(slice)
	fmt.Println(slice)

	// Pass the slice as a pointer. We can modify the original slice.
	squareSlicePtr(&slice)
	fmt.Println(slice)
}
