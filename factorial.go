package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		fmt.Println(1)
		return 1
	} else {
		fmt.Println(number)
		return factorial(number-1) * number
	}
}

// func main() {
// 	fmt.Println(factorial(4))
// }

// note

// (recursion 1)
// number = 4
// 4 is not equal to 1 so jump to else
// factorial(3) * number (recursion 2)
// it is not return yet because we need to finish factorial first

// number = 3
// 3 is not equal to 1 so jump to else
// factorial(2) * number (recursion 3)
// it is not return yet because we need to finish factorial first

// number = 2
// 2 is not equal to 1 so jump to else
// factorial(1) * number (recursion 4)
// it is not return yet because we need to finish factorial first

// number = 1
// recursion (5)
// 1 is equal to 1 so return 1
// we continue recursion 4

// number = 2
// factorial(1) = 1 so factorial(2) = factorial(1) * 2
// factorial(2) = 1 * 2 = 2

// jump to recursion 3
// number = 3
// factorial(2) = 2 so factorial(3) = factorial(2) * 3
// factorial(3) = 2 * 3 = 6

// jump to recursion 2
// number = 4
// factorial(3) is 6 so factorial(4) = factorial(3) * 4
// factorial(4) = 6 * 4 = 24

// so the result for factorial(4) is 24
