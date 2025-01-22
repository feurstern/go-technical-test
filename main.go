package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func calculateLetterAndNumber(i string) (l int, n int) {
	for _, x := range i {
		if unicode.IsLetter(x) {
			l++
		} else if unicode.IsDigit(x) {
			n++
		}
	}

	return
}

func main() {

	userInput := bufio.NewReader(os.Stdin)
	fmt.Println("enter any string(letter and number only!:")

	sentence, err := userInput.ReadString('\n')
	if err != nil {
		return
	}

	sentence = sentence[:len(sentence)-1]

	letters, numbers := calculateLetterAndNumber(sentence)
	fmt.Println("Total letters:", letters)
	fmt.Println("Total numbers:", numbers)
}
