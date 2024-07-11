package main

import (
	"flag"
	"fmt"
)

func main() {
	outputFlag := flag.String("output", "output.txt", "Output file name")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if args[0] == "" {
		return
	}
	if args[0] == "\\n" {
		fmt.Println()
		return
	}
	var banner, input, output string
	output = *outputFlag
}
