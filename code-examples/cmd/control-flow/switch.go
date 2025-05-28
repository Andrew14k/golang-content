package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	currentTime := time.Now().Hour()
	timeOfDay(currentTime)

	var letter string
	fmt.Print("input letter: ")
	fmt.Scanln(&letter)
	isVowel(letter)

}

func timeOfDay(currentTime int) {
	switch {
	case currentTime >= 0 && currentTime < 6:
		fmt.Println("It's too early, go back to sleep!")
	case currentTime >= 6 && currentTime < 12:
		fmt.Println("Good morning!")
	case currentTime >= 12 && currentTime < 13:
		fmt.Println("Out to lunch!")
	case currentTime >= 13 && currentTime < 18:
		fmt.Println("Good afternoon!")
	case currentTime >= 18 && currentTime < 22:
		fmt.Println("Good evening!")
	case currentTime >= 22 && currentTime < 24:
		fmt.Println("Good night!")
	default:
		fmt.Println("this is not a valid time")
	}
}

func isVowel(letter string) {
	letter = strings.ToLower(letter)
	switch {
	case len(letter) != 1 || letter < "a" || letter > "z":
		fmt.Println("this is not a valid letter")
	case letter == "a" || letter == "e" || letter == "i" || letter == "o" || letter == "u":
		fmt.Println("this is a vowel")
	default:
		fmt.Println("this is a consonant")
	}
}
