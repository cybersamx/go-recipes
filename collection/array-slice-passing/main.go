package main

import "fmt"

func squareArrayVal(array [3]int) {
	fmt.Printf("address of array in squareArrayVal: %p\n", &array)
	for i, v := range array {
		array[i] = v * v
	}
}
func squareArrayPtr(array *[3]int) {
	fmt.Printf("address of array in squareArrayPtr: %p\n", array)
	for i, v := range *array {
		array[i] = v * v
	}
}
func squareSliceVal(slice []int) {
	fmt.Printf("address of array in squareSliceVal: %p\n", &slice)
	for i, v := range slice {
		slice[i] = v * v
	}
}
func squareSlicePtr(slice *[]int) {
	fmt.Printf("address of array in squareSlicePtr: %p\n", slice)
	for i, v := range *slice {
		(*slice)[i] = v * v
	}
}
func main() {
	array := [3]int{1, 2, 3}

	fmt.Println("--- Array ---")
	fmt.Printf("Address of array: %p\n", &array)
	fmt.Println("original:              ", array)

	// Pass the array as a value. We get a copy of the array.
	// We can't modify the original array.
	squareArrayVal(array)
	fmt.Println("first pass by value:   ", array)

	// Pass the array as a pointer. We can modify the original array.
	squareArrayPtr(&array)
	fmt.Println("second pass by pointer:", array)

	slice := []int{1, 2, 3}

	fmt.Println("--- Slice ---")
	fmt.Printf("Address of slice-basics: %p\n", &slice)
	fmt.Println("original:              ", slice)

	// Pass the slice-basics as a value. We get a copy of the slice-basics variable, which references
	// the original data structure. We can modify the original slice-basics.
	squareSliceVal(slice)
	fmt.Println("first pass by value:   ", slice)

	// Pass the slice-basics as a pointer. We can modify the original slice-basics.
	squareSlicePtr(&slice)
	fmt.Println("second pass by pointer:", slice)
}
