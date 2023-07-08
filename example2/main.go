package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input int

	fmt.Print("Enter Number: ")
	if scanner.Scan() {
		inputStr := scanner.Text()
		var err error
		input, err = strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Invalid input.")
			return
		}
	}

	switch {
	case input%15 == 0:
		fmt.Println("FizzBuzz")
	case input%3 == 0:
		fmt.Println("Fizz")
	case input%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(input)
	}
}
