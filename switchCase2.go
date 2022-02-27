package main

import "fmt"

func time() int {
	return 6
}

func main() {
	switch work := time(); { // if the variable declared here, it will be local to this only
	case work < 6:
		fmt.Println("more focus")
	case work < 10:
		fmt.Println("work in focused burst to use time effectively")
	default:
		fmt.Println("other things dont worth your time")
	}

	fmt.Println(work) //can't access work

}
