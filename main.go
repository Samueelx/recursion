package main

import "fmt"

func factorial(num uint32) uint32 {
	var result uint32
	if num == 0 {
		return 1
	}

	result = num * factorial(num-1)
	return result
}

func main() {
	var input uint32
	fmt.Printf("Give me a number: \n")
	fmt.Scanf("%d", &input)
	fact := factorial(input)
	fmt.Printf("The factorial of the number is: %d\n", fact)
}
