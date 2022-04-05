package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	var numbers []float64

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}
