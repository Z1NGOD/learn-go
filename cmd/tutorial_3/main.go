package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "my new string"
	sendSting(printValue)

	var result, remainder, err = inDivision(5, 2)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("this will return result %v, and this remainder %v", result, remainder)
}

func sendSting(newString string) {
	fmt.Println(newString)
}

func inDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("You cant devide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
