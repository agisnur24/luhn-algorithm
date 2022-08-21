package main

import "fmt"

func main() {
	CardVerifier([]int{4, 5, 5, 6, 7, 3, 7, 5, 8, 6, 8, 9, 9, 8, 5, 5})
}

func CardVerifier(n []int) {
	var doublingSecondDigit, oddDigit, sum, digit int

	for i := 0; i < len(n); i++ {
		if i%2 == 0 {
			digit = n[len(n)-i-2] * 2
			if digit > 9 {
				digit -= 9
			}
			doublingSecondDigit += digit
		} else {
			oddDigit += n[i]
		}
	}
	fmt.Println("Sum doubling second digit from right to left = ", doublingSecondDigit)
	fmt.Println("Sum odd digit from right to left = ", oddDigit)

	sum = doublingSecondDigit + oddDigit

	fmt.Println("Sum step 1 & step 2 = ", sum)

	if sum%10 == 0 {
		fmt.Println("Status : Valid")
	} else {
		fmt.Println("Status : Invalid")
	}
}
