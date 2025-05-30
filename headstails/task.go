package main

import (
	"fmt"
	"math/rand"
)

// type function prompting string input and returns the value (heads/tails) with a 'quit' flag
type StringInput func(prompt string) (string, bool)

// type function prompting string input and returns the value (points num) with a 'quit' flag
type IntInput func(prompt string) (int, bool)

// ScoreCounter closure that tracks and updates players scores
func ScoreCounter() func(correct int, points int) int {
	score := 0

	return func(correct int, points int) int {
		if correct == 1 {
			score += points
		} else if correct == 0 {
			score -= points
		}
		return score
	}
}

// PlayerTurn handles the process for a players turn, selcting points and the turn itself
// getScore, getGuess, and getPoints taken as type aliases for the anonymous and type functions used
func PlayerTurn(player string, getGuess StringInput, getPoints IntInput, getScore func(int, int) int, flip string) (bool, int) {
	var correct int
	//Sprintf formats string to be passed to function, Printf prints directly to console
	points, quit := getPoints(fmt.Sprintf("%s, how many points do you want to play for? (type q to exit) ", player))
	if quit {
		return true, 0
	}

	guess, quit := getGuess(fmt.Sprintf("%s, heads or tails: (type q to exit) ", player))
	if quit {
		return true, 0
	}

	if guess == flip {
		correct = 1
		fmt.Printf("Correct! You got some points! %s Score: %d\n", player, getScore(correct, points))
		fmt.Println(" ")
	} else {
		correct = 0
		fmt.Printf("InCorrect! You lost some points :( %s Score: %d\n", player, getScore(correct, points))
		fmt.Println(" ")
	}
	return false, 0
}

func main() {
	coinFlip := [2]string{"heads", "tails"}

	//declare score counters
	playerA := ScoreCounter()
	playerB := ScoreCounter()

	fmt.Println("Welcome to the Heads or Tails game!")
	fmt.Println(" ")

	//anonymous function that reads a string input and checks for quit value
	getStr := StringInput(func(prompt string) (string, bool) {
		var input string
		fmt.Print(prompt)
		fmt.Scanln(&input)
		if input == "q" || input == "Q" {
			return "", true
		}
		return input, false
	})

	// anonymous function that reads an integer and checks for input error (quit value) - NEED TO FIX THIS PART
	getInt := IntInput(func(prompt string) (int, bool) {
		var input int
		fmt.Print(prompt)
		_, err := fmt.Scanln(&input)
		if err != nil {
			return 0, true
		}
		return input, false
	})

	for {
		flip := coinFlip[rand.Intn(2)] //selects heads/tails randomly

		//player A turn
		quit, _ := PlayerTurn("Player A", getStr, getInt, playerA, flip)
		if quit {
			break
		}

		//playerB turn
		quit, _ = PlayerTurn("Player B", getStr, getInt, playerB, flip)
		if quit {
			break
		}
	}
	fmt.Println(" ")
	fmt.Println("Game Over! Thanks for playing!")
	fmt.Printf("Player A final score: %d\n", playerA(0, 0))
	fmt.Printf("Player B final score: %d\n", playerB(0, 0))
	fmt.Println(" ")
}

// Concept:
// tracks scores of game between 2 players - score/penalise (gain/lose)	DONE
// turn based - number of points to be gained or lose can be selected at the start of each game DONE

// Rules:
// step 1 - point gained or lost each turn DONE
// step 2 - points gained or lost can be selected - can be changed throughout game DONE

// Reqs:
// Track scores for each player DONE
// Allow user to select scoring system DONE
// consider anonymous funcs, callbacks, closures DONE

// Game - heads or tails DONE

// FUTURE WORK:
// modify the input anonymous function to account for incorrect inputs that are not numbers or quit
// modify the string anonymous function to account for incorrect inputs that are not heads/tails or quit
// caps on point values?
// change what happens if score is to go < 0 ?
