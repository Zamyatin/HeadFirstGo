package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := argsToFloats(os.Args[1:])

	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

func average(numbers ...float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func argsToFloats(arguments []string) []float64 {
	var numbers []float64

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	return numbers
}
