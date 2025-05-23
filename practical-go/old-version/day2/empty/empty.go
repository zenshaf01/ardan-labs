package main

import "fmt"

// YOu can have a generic interface field
type Number interface {
	int | float64
}

func main() {
	// This is an empty interface. Any means i can become any type.
	// DONT USE empty interfaces or any
	var i any

	i = 7
	fmt.Println("i can be anything 1: ", i)

	i = "Seven"
	fmt.Println("i can be anything 2: ", i)

	s := i.(string) // Type assertion
	fmt.Println("s: ", s)

	n, ok := i.(int)
	if ok {
		fmt.Println("n: ", n)
	} else {
		fmt.Println("Not an int")
	}

	// Below is the type switch
	switch i.(type) {
	case int:
		fmt.Println("Is an int")
	}

	fmt.Println("Max: ", max([]int{1, 2, 3, 4, 5, 6}))
}

/*
	This is generic function which accepts a list of either ints or float64 and returns a one of the type.

	You should write a generic function when the method logic is exactly the same for two or more types.
	YOu should also do that when you are trying to writ6e a data structure which can take / house more than one type.

	Generics started with go 1.18
*/
func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]

	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
