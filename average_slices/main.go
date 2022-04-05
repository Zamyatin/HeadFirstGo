package main

import (
	"fmt"

	"github.com/zamyatin/go-averages/datafile"
)

func main() {
	fmt.Println(datafile.GetFloats("data.txt"))
}
