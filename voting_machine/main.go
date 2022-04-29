package main

import (
	"fmt"
	"log"

	"github.com/zamyatin/HeadFirstGo/average_slices/datafile"
)

func main() {
	votes, err := datafile.GetStrings("voteData.txt")

	if err != nil {
		log.Fatal(err)
	}

	var voteCount = make(map[string]int)

	for _, name := range votes {
		voteCount[name]++
	}

	for k, v := range voteCount {
		fmt.Printf("%s: %d\n", k, v)
	}
}
