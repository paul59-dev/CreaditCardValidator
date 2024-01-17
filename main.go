package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var (
		numberCard   string
		multiply     bool
		splitNumbers []string
		sum          int
	)

	numberCard = "12345"
	multiply = true

	// algo Luhn
	splitNumbers = strings.Split(numberCard, "")
	fmt.Println(splitNumbers)

	// display all number
	for _, splitNumber := range splitNumbers {
		if multiply {
			number, err := strconv.Atoi(splitNumber) // convert string numbre to interger number
			if err != nil {
				fmt.Println("Erreur de conversion: ", err)
				return
			}

			number = number * 2
			multiply = false

			if number >= 10 {
				// Convert number to string
				numberToStrings := strconv.Itoa(number)
				// split number
				splitNumbers = strings.Split(numberToStrings, "")
				// each value to convert to int
				for _, splitNumber := range splitNumbers {
					numberAdd, err := strconv.Atoi(splitNumber)
					if err != nil {
						fmt.Println("Erreur de conversion: ", err)
						return
					}

					sum += numberAdd
				}
				fmt.Println(sum)
			} else {
				fmt.Println(number)
			}

		} else {
			multiply = true
			fmt.Println(splitNumber)
		}
	}
}
