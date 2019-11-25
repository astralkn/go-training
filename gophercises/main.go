package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "math_problems.csv", "Path to cvs problems file")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	check(err, fmt.Sprintf("Failes to open the CSV file %s\n", *csvFilename))
	fmt.Println(file.Name())
	f := csv.Reader{}
	lines, err := f.Read()
	check(err, "Could not read the csv file")
	for line := range lines {
		fmt.Println(line)
	}
}

func check(err error, str string) {
	if err != nil {
		fmt.Println(str)
		panic(err)
	}
}
