// numberGuess game
// the game generates a random number between 1 and 100.
// the user must try to guess the number

package main

import (
	"math/rand"
	"time"
	"bufio"
	"os"
	"fmt"
	"strconv"
	"regexp"
)

func main()  {
	fmt.Println(setupGame())
	gamePlay(generateRandomNumber(1,100), 0)
	os.Exit(0)
}

// error handler
func errCheck(err error, message string)  {
	if err != nil {
		panic(fmt.Sprintf(message + "%v\n", err))
	}
}

// returns the a string with the correct guess and the number of attempts the user has made
func gameWon(randomNumber, attempts  int) string{
	return fmt.Sprintf("You Guessed it! %v is correct.\nYou made %v attempts.\n", randomNumber, attempts)
}

// generate a random number within a given range
func generateRandomNumber(min, max int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}

// returns a hint to the user based upon their guess and the random number
func getHint(random, guess int) string  {
	if random > guess {
		return "try a little higher\n"
	}
	return  "try a little lower\n"
}

// retrieves the user input from the console
func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		return scanner.Text()
	}
	errCheck(scanner.Err(), "Error Getting User input")

	return scanner.Text()
}

// checks to make sure the input from user is of proper type and length
func isValidInput(userInput string) bool {
	validInt := regexp.MustCompile(`^[0-9]{1,2}$`)
	return validInt.MatchString(userInput)
}

// provides the intro text for game, may in the future take an input tot declare number range
func setupGame() string  {
	return fmt.Sprintln("\n\nGuess that number!!\nTry and find the random number between 1 and 100.\n\n ---------------------------------------")
	// could get user input to setup an custom range of numbers
}

// checks for exit input from user
func exitCheck(input string)  {
	if input == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	}
}

// -----------------------------------------

// a recursive function which controls the flow of the game
func gamePlay(randomNumber, attempts int) {

	fmt.Println("Enter your guess: ")

	userInput := getUserInput()

	exitCheck(userInput)

	if b := isValidInput(userInput); b != true{
		fmt.Printf("\n %v, is an invalid input\n Try again!", userInput)
		gamePlay(randomNumber, attempts)
	}

	guess, err := strconv.Atoi(userInput)
	errCheck(err, "error parsing user input")

	if randomNumber != guess {
		fmt.Printf("Nope, %s", getHint(randomNumber, guess))
		attempts++
		gamePlay(randomNumber,attempts)

	} else {
		attempts++
		fmt.Println(gameWon(randomNumber, attempts))
	}

}
