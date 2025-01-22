package problems

import "fmt"

func NUmber62() {
	var num int

	result := 1
	fmt.Println("enter the value")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("not a valid format!")
		return
	}

	for i := 0; i < num; i++ {
		result *= i + 1
	}

	fmt.Printf("Result  %d", result)

}
