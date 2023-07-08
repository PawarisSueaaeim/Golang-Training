package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	method := args[1]
	input := args[2:]

	numberInput := make([]int, len(input))
	output := 0

	for i, str := range input {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invilid input number")
		} else {
			numberInput[i] = num
		}
	}

	switch method {
	case "all":
		for _, input := range numberInput {
			output += input
		}
	case "odd":
		for _, input := range numberInput {
			if input%2 == 1 {
				output += input
			}
		}
	case "even":
		for _, input := range numberInput {
			if input%2 == 0 {
				output += input
			}
		}
	}

	fmt.Println("method: ", method)
	fmt.Println("number: ", input)
	fmt.Println("numberInput: ", numberInput)
	fmt.Println("output: ", output)
}
