package main

import "fmt"

func Intersection(arr1, arr2 []int) []int {
	test := make(map[int]bool)
	// making a hash table
	for _, value := range arr1 {
		test[value] = true
	}

	var result []int
	for _, value := range arr2 {
		// checking for values in arr2 exists in arr1 hashtaable or not
		// looking for a vlue in hash table takes constant amount of time
		if _, ok := test[value]; ok {
			// adding matched values to a slice
			result = append(result, value)
		}
	}
	return result

}

func main() {
	arr1 := []int{2, 4, 6, 9, 19}
	arr2 := []int{5, 0, 13, 8, 9, 2}
	fmt.Println(Intersection(arr1, arr2))
}
