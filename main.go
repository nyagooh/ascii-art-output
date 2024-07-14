package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	art "output/output"
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
	if len(args) == 1 && flag.NFlag() < 1 && !strings.HasSuffix(args[0], ".txt") {
		input = os.Args[1]
		banner = "standard"
		result := art.ProcessFile(banner, input)
		fmt.Println(result)
	} else if len(args) == 1 && flag.NFlag() == 1 {
		input = args[0]
		output = os.Args[1][9:]
		if !strings.HasSuffix(output, ".txt") {
			log.Fatal("wrong file extension: file must have a '.txt' extension")
		}
		art.LessArguments(output, input)

		//fs project incooperated here.
	} else if len(args) == 2 && flag.NFlag() < 1 {
		input = args[0]
		banner = args[1]
		if strings.HasSuffix(banner, ".txt") {
			banner = strings.ToLower(strings.TrimSuffix(banner, ".txt"))
		} else {
			banner = strings.ToLower(banner)
		}
		result := art.ProcessFile(banner, input)
		fmt.Println(result)

		//output incooperated
	} else if len(args) == 2 && flag.NFlag() == 1 {
		input = args[0]
		art.Ascii_Output(output, input, args)
	}
}
