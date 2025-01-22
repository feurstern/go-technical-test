package problems

import "fmt"

func NUmber64() {

	var n [50]float64
	var total int
	fmt.Print("Enter number of index: ")

	fmt.Scanln(&total)
	for i := 0; i < total; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scan(&n[i])
	}
	for i := 1; i < total; i++ {
		if n[0] < n[i] {
			n[0] = n[i]
		}

	}

	fmt.Print("The highest number is : ", n[0])
}
