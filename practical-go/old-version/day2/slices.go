package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		Slices are akin to arrays in other languages.
		It is a contiguous block of memory in RAM
		Declares a slice of ints.
		s being declared means that an empty slice is being created in memory
	*/
	var s []int
	// you can compare a slice with a nil. The only comparison that can be done on a slice with ==
	// This is fast
	fmt.Println("len", len(s)) // len is nil safe. meaning you can call len on a nil slice.
	if s == nil {
		fmt.Println("slice is nil")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // This is getting a slice of s2. it goes fropm index 1 to index 3
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100]) // This will generate an out of range panic

	s3 = append(s3, 100)

	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2) // s2 is changed as well

	/*
		below is creating a slice of ints with 0 length and 0 capacity
	*/
	var s4 []int
	// s4 := make([]int, 11)
	for i := 0; i < 10; i++ {
		s4 = appendInt(s4, i)
	}

	fmt.Println("s4", s4, len(s4), cap(s4))

	concat([]string{"A", "B"}, []string{"C", "D", "E"})

	vs := []float64{2, 1, 4, 3}
	v, err := median(nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("vs: ", v, vs)
}

/*
Example implementation of the append function in go
*/
func appendInt(s []int, v int) []int {
	// We get the next position to put the element in
	i := len(s)

	// We check if the underlying array has enough space
	if len(s) < cap(s) {
		// if there is space
		// create a new slice from start to length of s plus 1
		fmt.Println("has capacity")
		s = s[:len(s)+1]
	} else {
		fmt.Println("reallocating")
		// if there is no space
		// create a new slice with double the capacity / size
		s2 := make([]int, 2*len(s)+1)
		// copy elements into s2
		copy(s2, s)
		// assign s2 from start to finish to s
		s = s2[:len(s)+1]
	}
	// add v at the target position
	s[i] = v
	return s
}

func concat(s1, s2 []string) []string {
	// no for loops

	// create a new slice with double the capacity of s1 and s2
	// copy both to new slice

	totalLength := len(s1) + len(s2)
	s3 := make([]string, totalLength)
	copy(s3, s1)
	copy(s3[len(s1):], s2)

	fmt.Println("S3: ", s3)
	return s3
}

/*
Everything in go is pass by value meaning go creates a copy of the arguyment and passes it into the function
Slice is alos a copy BUT the slice holds a pointer inside it over the underlying array.
*/
func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// sort the values
	/*
		The below sort i actually sorting the underlying array even if the values slice
		is a copy of the vs variable we declared in the caller function.

		This is because the slice of values is a struct which is actually holding a pointer
		over the underlying array. So no matter if go creates a copy of the slice, the slice itself,
		holds a pointer over the array.values
		So whenever you are working with slices, make sure you understand what it is doing to avoid mutating data that you dont want to

		to fix the below just make a new slice and perform the median calculation on the new slice
	*/
	// sort.Float64s(values) // This is mutating the original vs slice
	nums := make([]float64, len(values))
	// Copy in order to not change the original values
	copy(nums, values)
	sort.Float64s(nums)

	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	/*
		constants type is defined at the place of usage.
		note that n is being used to divide with a float64.
		It would also work if the valus was an array of integers
	*/
	const n = 2
	v := (nums[i-1] + nums[i]) / n
	return v, nil
}
