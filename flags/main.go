package main

import (
	"flag"
	"fmt"
)

/*
Example flags
*/
func main() {
	fmt.Println("Staring the application...")
	output := flag.Bool("output", false, "Should there be an output?")
	input := flag.String("input", "file.csv", "The path to the input file")
	flag.Parse()
	fmt.Println(*output)
	fmt.Println(*input)
}
