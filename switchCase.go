package main

import "fmt"

const (
	normal       = 60
	grinding     = 80
	hardcoreMode = 100
)

func time() string {
	return "all in"
}

func main() {
	work := time()
	switch work {
	case "half work":
		fmt.Println("more focus")
	case "all in":
		fmt.Println("think effectively")
	default:
		fmt.Println("dont let your young you dreams down")
	}

	mode := grinding
	switch { // dont provide expression if u want to use conditions
	case mode < 50:
		fmt.Println("nothing wont come out of this")
	case mode < 70:
		fmt.Println("some more")
	case mode < 90:
		fmt.Println("continue grinding")
	default:
		fmt.Println("No time to spend on toher things, expcept one thing")
	}

}
