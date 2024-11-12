package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is a simple wordle app")
	answer := ""
	firstGuess := inputWord(answer)
	matchedGuess := ""
	for j := 0; j < len(answer); j++ {
		for i := 0; i < len(firstGuess); i++ {
			if string(firstGuess[i]) == string(answer[i]) {
				matchedGuess += string(firstGuess[i])
			} else {
				matchedGuess += "_"
			}
		}
		if matchedGuess == answer {
			fmt.Printf("You got it!!, the word is %v", answer)
			return
		} else {
			fmt.Println(matchedGuess)
			matchedGuess = ""
			firstGuess = inputWord(answer)
		}
	}

}

func inputWord(a string) string {
	word := bufio.NewReader(os.Stdin)

	fmt.Printf("Guess your word (%v character): ", len(a))
	guess, _ := word.ReadString('\n')
	guess = strings.TrimSpace(guess)
	return guess
}
