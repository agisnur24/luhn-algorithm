package main

import "fmt"

func main() {
	CardVerifier([]int{4, 5, 5, 6, 7, 3, 7, 5, 8, 6, 8, 9, 9, 8, 5, 5})
}

func CardVerifier(n []int) {
	var dsd, od, sum, d int

	for i := 0; i < len(n); i++ {
		if i%2 == 0 {
			d = n[len(n)-i-2] * 2
			if d > 9 {
				d -= 9
			}
			dsd = dsd + d
		} else {
			od += n[i]
		}
	}
	fmt.Println("Sum doubling sec digit from right-left = ", dsd)
	fmt.Println("Sum odd digit from right-left = ", od)

	sum = dsd + od

	fmt.Println("Sum step 1 & step 2 = ", sum)

	if sum%10 == 0 {
		fmt.Println("Status : Valid")
	} else {
		fmt.Println("Status : Invalid")
	}
}
