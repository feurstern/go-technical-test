package problems

import (
	"fmt"
	"math/rand"
)

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func generateRandomNumbers(n int, result []int) []int {

	if len(result) == n {
		return result
	}

	num := rand.Intn(n) + 1

	for contains(result, num) {
		num = rand.Intn(n) + 1
	}

	result = append(result, num)

	return generateRandomNumbers(n, result)

}

func Number4() {
	var number int

	fmt.Println("enter the value")
	_, err := fmt.Scan(&number)

	if err != nil {
		return
	}
	result := generateRandomNumbers(number, []int{})

	for i, num := range result {
		if i == len(result)-1 {
			fmt.Print(num)
		} else {
			fmt.Print(num, ", ")
		}
	}
}
