package main

import "fmt"

func main() {
	slice := []string{"kalios", "lets", "get", "this", "done"}
	fmt.Println(slice[0])
	updateSlice(slice)
	
	// value changes because it is a reference type
	// under the hood it is pointing to the array 
	fmt.Println(slice[0])
}

func updateSlice(slice []string) {
	slice[0] = "alex"
}
