package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var (
		numberCard      string
		multiply        bool
		splitNumbers    []string
		sum             int
		sliceNumberSums []int
		sumOfAll        int
		numberOfControl int
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
				sliceNumberSums = append(sliceNumberSums, sum)
				fmt.Println(sum)
			} else {
				sliceNumberSums = append(sliceNumberSums, number)
				fmt.Println(number)
			}

		} else {
			multiply = true
			number, err := strconv.Atoi(splitNumber) // convert string numbre to interger number
			if err != nil {
				fmt.Println("Erreur de conversion: ", err)
				return
			}
			sliceNumberSums = append(sliceNumberSums, number)
			fmt.Println(splitNumber)
		}
	}

	fmt.Println(sliceNumberSums)

	sumOfAll = 0
	for i := 0; i < len(sliceNumberSums); i++ {
		sumOfAll += sliceNumberSums[i]
	}

	fmt.Println("somme de tout les chiffre: ", sumOfAll)

	numberOfControl = 10 - (sumOfAll % 10)
	fmt.Println(numberOfControl)
	sliceNumberSums = append(sliceNumberSums, numberOfControl)
	fmt.Println(sliceNumberSums)
}
