package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"virtual-file-system/internal/actions"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		text := scanner.Text()
		args := strings.Split(text, " ")
		f := &actions.Factory{}
		if act := f.CreateAction(args); !act.Exec(args) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
