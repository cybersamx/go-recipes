# Passing Array to Functions

<<<<<<< Updated upstream
In Go, arrays are passed to a function as values. So the receiving function gets a copy of the original array. This can be a problem if we need to change the array. To modify the original array, we need to pass the array as a reference or use a slice.

A slice, on the hand, isn't a data structure but rather a variable that references the original data structure. So whether you pass a slice (as a value or reference) to a function, only a copy of that variable is made, which still points to the original data structure.
=======
In Go, an array is passed to a function as a value. So the receiving function gets a copy of the original array. This can be a problem if we need to change the array. To modify the original array, we need to pass the array as a pointer or pass the values as a slice.

A slice, on the hand, is passed to a function not as a data structure but rather as a variable that references the original data structure. So whether you passes a slice as a value or reference to a function. Only a copy of that variable is made, the copy still points to the original data.
>>>>>>> Stashed changes

> **Notes**
>
> It's interesting to note that a slice has a different set of rules when it being passed to a function. The receiving function gets a copy of the original slice, but it still references the original array. This is due to the fact that slice is a struct consisting of a pointer to the array, length of the segment, and its capacity. See [Go Blog: Slice Intro](https://blog.golang.org/slices-intro) for details.

## Setup

1. Run the program.

   ```bash
   $ make run
   ```

## Reference

* [Go Blog: Slice Intro](https://blog.golang.org/slices-intro)
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)

