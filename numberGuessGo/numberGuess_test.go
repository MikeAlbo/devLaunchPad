// testing for numberGuess.go
package main

import (
	"testing"
	"fmt"
)

// number generator
func TestRandomNumberGeneratorToReturnNewValues(t *testing.T)  {

	var actualReturnedValues  []int

	for i:=0; i < 5; i++ {
		actualReturnedValues = append(actualReturnedValues, generateRandomNumber(1, 100) )
	}

	for j := 0; j < len(actualReturnedValues); j++ {
		if j < len(actualReturnedValues) -1 {
			if actualReturnedValues[j] == actualReturnedValues[j+1] {
				fmt.Println(actualReturnedValues)
				t.Fatalf("Random Number Generator is generating identical ints.")
			}

		} else {
			if actualReturnedValues[j] == actualReturnedValues[0] {
				fmt.Println(actualReturnedValues)
				t.Fatalf("Random Number Generator is generating identical ints.")
			}
		}
	}

	fmt.Println(actualReturnedValues)


}

func TestRandomNumberGeneratorToReturnValuesWithinRange(t *testing.T){
	for i:=0; i < 100; i++ {
		if randomNumber := generateRandomNumber(1,100); randomNumber < 1 || randomNumber > 100 {
			t.Fatalf("randomNumberGenerator returned the value %v, which is outside the scope of the given range", randomNumber)
		}
	}
}

// hint feedback
func TestGetHintToReturnCorrectString(t *testing.T)  {
	r,g := 20,25
	if getHint(r, g) != "try a little lower\n" {
		t.Fatalf("random: %v, and guess: %v, should return, try a little lower ", r,g )
	} else if getHint(g,r) != "try a little higher\n" {
		t.Fatalf("random: %v, and guess: %v, should return, try a little higher ", g,r )
	}
}

func TestGameWonScenarioToReturnCorrectValues(t *testing.T){
	randomNumber, attempts  := 65,4

	expectedOutput := fmt.Sprintf("You Guessed it! %v is correct.\nYou made %v attempts.\n", randomNumber, attempts)

	if gameWon(randomNumber, attempts) != expectedOutput {
		t.Fatalf("The Game Won scenrio is not producing the correct output.\n-- %s ", expectedOutput)
	}
}

func TestIsValidInputReturnsTrueForInts(t *testing.T){
	input := []string{"1", "2", "3", "10","22", "50"}
	for _, v := range input {
		if b:= isValidInput(v); b != true {
			t.Fatalf("Is valid input, does not recognize %v as a valid input", v)
		}
	}

}

func TestIsValidInputReturnsFalseForNonIntsAndNumbersGreaterThan99(t *testing.T){
	input := []string{"one", "two", "100", "1p", "p1"}
	for _, v := range input {
		if b:= isValidInput(v); b != false {
			t.Fatalf("Is valid input, should return false for %v", v)
		}
	}

}

