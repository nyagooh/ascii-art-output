package output

import (
	"fmt"
	"os"
)

func PrintToFile(banner, input string) {
	file, err := os.Create(banner)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	_, err = file.WriteString(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}

func newline(file *os.File) {
	_, err := file.WriteString("\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
