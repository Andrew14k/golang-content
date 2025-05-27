package main

import "fmt"

func main() { //dont want to clutter up main use
	var timeOfDay int = 21
	fmt.Println(greeting(timeOfDay))
}

// used refactor function on rightclick
func greeting(timeOfDay int) string {
	if timeOfDay >= 5 && timeOfDay < 12 {
		return "Good morning"
	} else if timeOfDay >= 12 && timeOfDay <= 18 {
		return "Good afternoon"
	} else if timeOfDay > 18 && timeOfDay <= 23 {
		return "Good evening"
	} else {
		return "Good evening"
	}
}
