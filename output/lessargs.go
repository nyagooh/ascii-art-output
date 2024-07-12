package output

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
This function will deal with cases where the flag is provided with the output
file which is not either of the banner files present but has an extension of txt
while the banner file to be read from isn't provided.
We pass in the desired output file name and the string input.
*/
func LessArguments(output, input string) {
	if len(output) < 4 {
		log.Fatal("Error: The .txt file name length size should be greater than 4")
	}
	if output == "shadow.txt" || output == "standard.txt" || output == "thinkertoy.txt" {
		log.Fatal("The name of the text file is the same as the banner file name. Please use a different file name.")
	}
	if strings.Count(output, ".") > 1 {
		log.Fatal("Usage: filename.txt")
	}
	file, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line != "" {
			result := ProcessFile("standard", line)
			WriteinFile(file, result)
		} else {
			newline(file)
		}
	}
}
