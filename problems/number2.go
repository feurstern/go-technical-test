package problems

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func Number2() {
	str := bufio.NewReader(os.Stdin)
	fmt.Println("enter any string")

	sentence, err := str.ReadString('\n')

	if err != nil {
		return
	}

	sentence = sentence[:len(sentence)-1]

	newStr := reverseString(sentence)
	fmt.Printf("the sentence you input: %s\n", sentence)

	fmt.Println("is it palindrome: ", newStr == sentence)
}
