package main

import (
	"flag"
	"fmt"
	"strings"
	"log"
	"os"
	art "output/functions"
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
		input = args[0]
		banner = "standard"
		result := art.ProcessFile(banner, input)
		fmt.Println(result)
		output = "output.txt"
	}else if len(args) == 1&& flag.NFlag() == 1{
		if !strings.HasPrefix(os.Args[1],"--output="){
			log.Fatal("wrong flag format:flag must have '--'")
		}
		output = strings.TrimPrefix(output, "--output=")
		if !strings.HasSuffix(output, ".txt") {
			log.Fatal("wrong file extension: file must have a '.txt' extension")
		}
		input = args[0]
		banner = "standard"
	}else if len(args) == 2 && flag.NFlag() < 1 && !strings.HasSuffix(args[0], ".txt") {
		input = args[0]
		if strings.HasSuffix(args[1], ".txt") {
			banner = strings.TrimSuffix(args[1], ".txt")
		} else {
			banner = strings.ToLower(args[1])
		}

		result := art.ProcessFile(banner, input)
		fmt.Println(result)
		output = "output.txt"
	} else if len(args) == 2 && flag.NFlag() == 1 {
		if strings.HasPrefix(args[1],".txt") {
			new := strings.TrimRight(args[1], ".txt")
			banner = strings.ToLower(new)
		} else {
			banner = strings.ToLower(args[1])
		}
		input = args[0]
		if !strings.HasPrefix(os.Args[1],"--output="){
			log.Fatal("wrong flag format:flag must have '--'")
		}
		output = strings.TrimPrefix(output, "--output=")
		if !strings.HasSuffix(output, ".txt") {
			log.Fatal("wrong file extension: file must have a '.txt' extension")
		}
	}
	if len(output) < 4 {
		log.Fatal("Error: The .txt file size should be greater than 4")
	}
	if output == "shadow.txt" || output == "standard.txt" || output == "thinkertoy.txt" {
		log.Fatal("The name of the text file is the same as the banner file name. Please use a different file name.")
	}
	if strings.Count(output, ".") > 1 {
		log.Fatal("Usage: filename.txt")
	}
	result := art.ProcessFile(banner, input)
}
