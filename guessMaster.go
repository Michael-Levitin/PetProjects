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
	fmt.Println("An answer to you guess will be given in a form of 'X bull/s, Y cow/s' ")
	fmt.Println("A 'bull' means that you guessed a digit at the right place")
	fmt.Println("A 'cow' means - there's a correct digit at a wrong place")
	fmt.Println("(All the digits in the sequence are different - example: 0389)")
	fmt.Println("(That's also apply to your guess)")
	startGame()
}

func startGame() {
	var letter string
	fmt.Println("Wanna play - Y/N?")

	fmt.Scan(&letter)
	switch letter {
	case "y":
		fallthrough
	case "Y":
		guessMaster()
	//case "n": fallthrough
	default:
		fmt.Println("Goodbye :)")
	}
}

func guessMaster() {
	var guess [4]int
	surrender := [4]int{0, 0, 0, 0}
	mistake := [4]int{1, 1, 1, 1}

	sequence := randomDigits()

	fmt.Println("Enter you guess in \"0389\" format or S to surrender") //
	for i := 1; ; i++ {
		//scanned, err := fmt.Scanf("%d %d %d %d", &guess[0], &guess[1], &guess[2], &guess[3])
		//scanned, err := fmt.Scan(&guess[0], &guess[1], &guess[2], &guess[3])
		guess = getGuess()
		// https://stackoverflow.com/questions/56103775/how-to-print-formatted-string-to-the-same-line-in-stdout-with-go
		if guess == surrender {
			fmt.Println("Sorry to hear that, better luck next time")
			fmt.Println("The answer is:", sequence, "\n")
			break
		} else if guess == mistake {
			fmt.Println("- Please enter 4 digits")
			continue
		} else if guess == sequence {
			fmt.Printf("Congratulations, You've Won in %d moves", i)
			break
		} else {
			checkGuess(guess, sequence)
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

func checkGuess(guess [4]int, sequence [4]int) {
	var bull int
	var cow int
	setGuess := make(map[int]bool)
	setSequence := make(map[int]bool)

	for i := 0; i < 4; i++ { // check unique digits
		setGuess[guess[i]] = true

	}
	if len(setGuess) != 4 {
		fmt.Println("Guess must contain unique digits")
		return
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

	return
}

func getGuess() [4]int {
	var line string
	var guess [4]int
	surrender := [4]int{0, 0, 0, 0}
	mistake := [4]int{1, 1, 1, 1}
	var split []string

	fmt.Scanln(&line)

	//in := bufio.NewReader(os.Stdin)
	//line,_ := in.ReadString('\n')
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
