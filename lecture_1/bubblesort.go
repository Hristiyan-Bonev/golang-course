package main

import (
	"fmt"
	"os"
)

// BubbleSort This function sorts exactly 5 integers (using bubble sort algorithm)
func BubbleSort(ints ...int) ([]int, error) {
	if len(ints) != 5 {
		return nil, fmt.Errorf("can sort exactly 5 numbers only")
	}

	tmp := make([]int, len(ints))

	// Copy to temp buffer to not mutate the provided parameters
	copy(tmp, ints)

	for i := len(ints); i > 0; i-- {
		for j := 1; j < i; j++ {
			if tmp[j-1] > tmp[j] {
				tmp[j], tmp[j-1] = tmp[j-1], tmp[j]
			}
		}
	}
	return tmp, nil
}

func main() {
	numsToSort := []int{-99, 2, 3, -4, 1}
	fmt.Printf("Sorting: %v\n", numsToSort)
	res, err := BubbleSort(numsToSort...)
	if err != nil {
		fmt.Printf("error while sorting: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", res)
	os.Exit(0)

}
