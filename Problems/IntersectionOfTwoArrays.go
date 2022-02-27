package main

import "fmt"

func Intersection(arr1, arr2 []int) []int {
	test := make(map[int]bool)

	for _, value := range arr1 {
		test[value] = true
	}

	var result []int
	for _, value := range arr2 {
		if _, ok := test[value]; ok {
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
