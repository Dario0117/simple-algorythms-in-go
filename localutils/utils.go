package localutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt ...
func Prompt(message string) string {
	var input string
	fmt.Printf(message + "> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

// PrintLineInsideMenu ...
func PrintLineInsideMenu(text string) {
	fmt.Printf("# %-51s #\n", text)
}

// PrintHeadFootPattern ...
func PrintHeadFootPattern() {
	fmt.Println(strings.Repeat("#", 55))
}
