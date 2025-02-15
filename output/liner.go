package output

// calculates mathematically the line number of each rune contained in the input string.
func FndLine(r rune) []int {
	lineNumbers := make([]int, 8)
	for i := 0; i < 8; i++ {
		lineNumbers[i] = (int(r-32) * 9) + i + 2
	}
	return lineNumbers
}

/*
connects the input string and the FndLine() function
*/
func ProcessLine(line string) []int {
	var result []int
	for _, char := range line {
		result = append(result, FndLine(char)...)
	}
	return result
}
