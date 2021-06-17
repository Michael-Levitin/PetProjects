package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("GuessMaster is a game, where you guess 4-digit sequence ")
	fmt.Println("An answer to your guess will be given in a form of 'X bull/s, Y cow/s' ")
	fmt.Println("A 'bull' means that you guessed a digit at the right place")
	fmt.Println("A 'cow' means - there's a correct digit at a wrong place")
	fmt.Println("(All the digits in the sequence are different - example: 0389)")
	fmt.Println("(That's also apply to your guess)")
	startGame()
}

func startGame() {
	var letter string
	fmt.Println("\nWanna play - Y/N? or should I guess your sequence - P")

	fmt.Scan(&letter)
	switch letter {
	case "y":
		fallthrough
	case "Y":
		guessMaster()
	case "p":
		fallthrough
	case "P":
		guessMasterPC()
	//case "n": fallthrough
	default:
		fmt.Println("Goodbye :)")
	}
}

func guessMasterPC() {
	//var s [][4]int
	var bulls int
	var cows int

	s := makeSliceOfAnswers()

	fmt.Println("Ready to guess your sequence. ")
	fmt.Println("Please write you response as 2 digits (bull/s cow/s) - '2 0' ")
	for {
		if len(s) == 1 {
			fmt.Println("Your number is ", s[0])
			break
		}
		if len(s) == 0 {
			fmt.Println("I surrender, you've probably made a mistake")
			break
		}

		rand.Seed(time.Now().UnixNano()) // for truly? random
		num := rand.Intn(len(s))
		currentGuess := s[num]
		fmt.Println("I'm guessing", currentGuess)

		for j := 0; ; j++ {
			if _, err := fmt.Scan(&bulls, &cows); err != nil {
				fmt.Println("I didn't understood that, can you rephrase?")
			} else {
				eliminateAnswers(currentGuess, &s, bulls, cows)
				break
			}

			if j > 2 {
				fmt.Println("Lets try other number")
				break
			}
		}
	}
	startGame()
}

func guessMaster() {
	var guess [4]int
	surrender := [4]int{0, 0, 0, 0}
	mistake := [4]int{1, 1, 1, 1}

	sequence := randomDigits()

	fmt.Println("Enter you guess in \"0389\" format or S to surrender") //
	for i := 1; ; i++ {
		guess = getGuess()
		if guess == surrender {
			fmt.Println("Sorry to hear that, better luck next time")
			fmt.Println("The answer is:", sequence)
			break
		} else if guess == mistake {
			fmt.Println("- Please enter 4 digits")
			continue
		} else if guess == sequence {
			fmt.Printf("Congratulations, You've Won in %d moves\n", i)
			break
		} else {
			bull, cow := checkGuess(guess, sequence)
			printGuessResponse(bull, cow)
		}
	}
	startGame()
}

func randomDigits() [4]int {
	set := make(map[int]bool)
	var array [4]int
	for i := 0; i < 4; {
		rand.Seed(time.Now().UnixNano()) // for truly? random
		num := rand.Intn(10)             //10 not included
		if exists := set[num]; !exists { // digits must be unique
			set[num] = true
			array[i] = num
			i++
		}
	}
	return array
}

func makeSliceOfAnswers() [][4]int {
	var s [][4]int
	index := 0
	s = make([][4]int, 5040)
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			if i == j {
				continue
			}
			for k := 0; k <= 9; k++ {
				if i == k || j == k {
					continue
				}
				for l := 0; l <= 9; l++ {
					if i == l || j == l || k == l {
						continue
					} else {
						s[index] = [4]int{i, j, k, l}
						//fmt.Println(s[index])
						index++
					}
				}
			}
		}
	}
	return s
}

func eliminateAnswers(guess [4]int, answers *[][4]int, bullsAnswer int, cowsAnswer int) {
	for i := 0; i < len(*answers); {
		bullsCompare, cowsCompare := checkGuess((*answers)[i], guess)
		if bullsCompare != bullsAnswer ||
			cowsCompare != cowsAnswer {
			*answers = append((*answers)[:i], (*answers)[i+1:]...)
		} else {
			i++ // advance only if not eliminating
		}
	} // if checkGuess doesn't give same bulls/cows - delete this answer
	//fmt.Println(len(*answers))
}

func checkGuess(guess [4]int, sequence [4]int) (int, int) {
	var bull int
	var cow int
	setGuess := make(map[int]bool)
	setSequence := make(map[int]bool)

	for i := 0; i < 4; i++ { // check unique digits
		setGuess[guess[i]] = true
	}
	if len(setGuess) != 4 {
		fmt.Println("Guess must contain unique digits")
		bull = -1
		cow = -1
		return bull, cow
	}

	for i := 0; i < 4; i++ { // populate sequence set
		setSequence[sequence[i]] = true
	}

	for i := 0; i < 4; i++ { // calculating guess
		if guess[i] == sequence[i] {
			bull++
		} else if exists := setSequence[guess[i]]; exists {
			cow++
		}
	}
	return bull, cow
}

func printGuessResponse(bull int, cow int) {
	if bull == -1 {
		return
	}

	fmt.Print(bull, " bull") // Formatting answer
	if bull != 1 {
		fmt.Print("s, ")
	} else {
		fmt.Print(", ")
	}
	fmt.Print(cow, " cow")
	if cow != 1 {
		fmt.Println("s")
	} else {
		fmt.Println("")
	}
}

func getGuess() [4]int {
	var line string
	var guess [4]int
	surrender := [4]int{0, 0, 0, 0}
	mistake := [4]int{1, 1, 1, 1}
	var split []string

	fmt.Scanln(&line)
	strings.Trim(line, "\n")

	if line == "S" || line == "s" { //len(line) == 1 &&(
		guess = surrender
	} else if len(line) == 4 {
		split = strings.Split(line, "")
		for i := range split {
			num, err := strconv.Atoi(split[i])
			if err != nil {
				guess = mistake
				break
			} else {
				guess[i] = num
			}
		}
	} else {
		guess = mistake
	}
	return guess
}
