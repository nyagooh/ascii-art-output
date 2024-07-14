package output

import (
	"fmt"
	"os"
	"strings"
)

func ValidFlag() {
	if strings.HasPrefix(os.Args[1], "--output") {
		if !strings.Contains(os.Args[1], "=") {
			fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]
			
EX: go run . --output=<fileName.txt> something "standard"`)
			os.Exit(0)
		}
	}
}
