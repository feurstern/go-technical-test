package problems

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Number5() {

	const inputFormat = "03:04:05 PM"
	userInput := bufio.NewReader(os.Stdin)
	fmt.Println("enthe the format")

	hourFormat, err := userInput.ReadString('\n')

	if err != nil {
		return
	}

	hourFormat = hourFormat[:len(hourFormat)-1]

	parsedTime, err := time.Parse(inputFormat, hourFormat)

	if err != nil {
		fmt.Println("err parsedtime")
		return
	}

	fmt.Printf("24 HOUR FORMAT %s:", parsedTime)

}
