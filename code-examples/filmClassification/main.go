package main

import "fmt"

func main() { //dont want to clutter up main use
	var age int = 11
	fmt.Println(GetAgeClassification(age))

}

func GetAgeClassification(age int) string {

	switch {
	case age < 12:
		return "U & PG films are available"

	case age < 15:
		return "U, PG & 12 films are available"

	case age < 18:
		return "U, PG, 12 & 15 films are available"
	default:
		return "All films are available"
	}

	/* // Same logic but in nested ifs for fun
	   if age < 12 {
	       return "U & PG films are available"
	   } else if age < 15 {
	       return "U, PG & 12 films are available"
	   } else if age < 18 {
	       return "U, PG, 12 & 15 films are available"
	   } else {
	       return "All films are available"
	   }
	*/

	/* // Original code
	   var result string

	   if age < 12 {

	       result = "U, PG & 12 films are available"
	   } else if age < 15 {

	       result = "U, PG, 12 & 15 films are available"
	   } else {

	       result = "All films are available"

	   }

	   return result
	*/
}
