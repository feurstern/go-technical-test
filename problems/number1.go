package problems

import "fmt"

func FirstQuestion() {
	var number, total int

	fmt.Println("Enter the valuee")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 0; i < number; i++ {
		fmt.Printf("number: %d \n", i+1)
		total += i + 1
	}

	fmt.Printf("total number of  file: %d\n", number)

	fmt.Printf("Sum of the number %d", total)
}
