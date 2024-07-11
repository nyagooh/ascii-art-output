package functions

import (
	"bufio"
	"log"
	"os"
)

/*
opens the required file,scans line by line and maps the line number to the strings found in the various
lines in the opened file.
*/
func Maps(n int, filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	files, err := file.Stat()
	if err != nil {
		log.Fatal("error calculating the file stat")

	}
	if !(files.Size() == 6624 || files.Size() == 7463 || files.Size() == 5558) {
		log.Fatal("the files are formatted in a way that they cant be accessed")
	}
	scanner := bufio.NewScanner(file)
	var count int
	var line, str string
	for scanner.Scan() {
		line = scanner.Text()
		count++
		ascMap := make(map[int]string)
		if count == n {
			ascMap[count] = line
			str = ascMap[n]
		}
	}
	if count == 0 {
		log.Fatal("empty banner file")
	}
	defer file.Close()
	return str
}
